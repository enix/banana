#! /usr/bin/env bash

auto="y"
org="enix"

function setIfEmpty {
  if [ -z "$3" ]; then
    export "$1"="$2"
  fi
}

function createCA {
  if [[ ${auto} == "y" ]]; then
    echo "=> using 'enix' as passphrase"
    openssl genrsa -des3 -passout pass:enix -out ../security/ca/$1-tmp.key 2048
    cat default.txt | openssl req -x509 -passin pass:enix -new -nodes -extensions v3_ca -key ../security/ca/$1-tmp.key -sha256 -days 2190 -out ../security/ca/$1.pem
  else
    openssl genrsa -des3 -out ../security/ca/$1-tmp.key 2048
    openssl req -x509 -new -nodes -extensions v3_ca -key ../security/ca/$1-tmp.key -sha256 -days 2190 -out ../security/ca/$1.pem
  fi

  openssl rsa -in ../security/ca/$1-tmp.key -out ../security/ca/$1.key -passin pass:enix -outform pem
  rm ../security/ca/$1-tmp.key
}

function createTLSCert {
  cert=$(curl -sSL \
    -H "X-Vault-Token: $VAULT_TOKEN" \
    -X POST \
    -d "{\"common_name\": \"$1\", \"ttl\": \"52500h\"}" \
    "$VAULT_ADDR/v1/pki/issue/$2")

  echo ${cert} | jq -r '.data.certificate' > ../security/tls/$1.pem
  echo ${cert} | jq -r '.data.private_key' > ../security/tls/$1.key
}

function init {
  mkdir -p ../security
  mkdir -p ../security/ca
  mkdir -p ../security/tls

  createCA ca
  createCA users-ca
  createCA agents-ca

  caCert=$(cat ../security/ca/ca.pem | sed -e 's/$/\\n/g' | tr -d '\n')$'\\n'
  caKey=$(cat ../security/ca/ca.key | sed -e 's/$/\\n/g' | tr -d '\n')
  usersCaCert=$(cat ../security/ca/users-ca.pem | sed -e 's/$/\\n/g' | tr -d '\n')$'\\n'
  usersCaKey=$(cat ../security/ca/users-ca.key | sed -e 's/$/\\n/g' | tr -d '\n')
  agentsCaCert=$(cat ../security/ca/agents-ca.pem | sed -e 's/$/\\n/g' | tr -d '\n')$'\\n'
  agentsCaKey=$(cat ../security/ca/agents-ca.key | sed -e 's/$/\\n/g' | tr -d '\n')

  vault secrets enable pki
  vault secrets tune -max-lease-ttl=53000h pki
  vault secrets enable --path=users-pki pki
  vault secrets tune -max-lease-ttl=53000h users-pki
  vault secrets enable --path=agents-pki pki
  vault secrets tune -max-lease-ttl=53000h agents-pki

  # Set root CA for TLS
  curl -sSL \
    -H "X-Vault-Token: $VAULT_TOKEN" \
    -X POST \
    -d "{\"pem_bundle\": \"${caCert}${caKey}\"}" \
    "$VAULT_ADDR/v1/pki/config/ca"

  # Create role for TLS
  curl -sSL \
    -H "X-Vault-Token: $VAULT_TOKEN" \
    -X POST \
    -d "{\"allow_any_name\": true, \"max_ttl\": \"52560h\", \"organization\": \"${org}\"}" \
    "$VAULT_ADDR/v1/pki/roles/${org}"

  createTLSCert banana.enix.io ${org}
  createTLSCert api.banana.enix.io ${org}

  # Set root CA for users
  curl -sSL \
    -H "X-Vault-Token: $VAULT_TOKEN" \
    -X POST \
    -d "{\"pem_bundle\": \"${usersCaCert}${usersCaKey}\"}" \
    "$VAULT_ADDR/v1/users-pki/config/ca"

  # Create role for users
  curl -sSL \
    -H "X-Vault-Token: $VAULT_TOKEN" \
    -X POST \
    -d "{\"allow_any_name\": true, \"max_ttl\": \"52560h\", \"organization\": \"${org}\"}" \
    "$VAULT_ADDR/v1/users-pki/roles/${org}"

  # Set root CA for agents
  curl -sSL \
    -H "X-Vault-Token: $VAULT_TOKEN" \
    -X POST \
    -d "{\"pem_bundle\": \"${agentsCaCert}${agentsCaKey}\"}" \
    "$VAULT_ADDR/v1/agents-pki/config/ca"

  # Create role for agents
  curl -sSL \
    -H "X-Vault-Token: $VAULT_TOKEN" \
    -X POST \
    -d "{\"allow_any_name\": true, \"max_ttl\": \"52560h\", \"organization\": \"${org}\"}" \
    "$VAULT_ADDR/v1/agents-pki/roles/${org}"

  cat ../security/ca/users-ca.pem > ../security/ca/trusted-ca.pem
  echo >> ../security/ca/trusted-ca.pem
  cat ../security/ca/agents-ca.pem >> ../security/ca/trusted-ca.pem
}

set -e

setIfEmpty VAULT_ADDR "http://127.0.0.1:7777" "${VAULT_ADDR}"
setIfEmpty VAULT_TOKEN "myroot" "${VAULT_TOKEN}"

cd $(dirname $0)

init $@

cd - > /dev/null

docker-compose restart proxy
echo $'\n=> ready'
