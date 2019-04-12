#! /usr/bin/env bash

if [[ $# -lt 1 ]] ; then
  echo "Usage: $0 <org>"
  exit 1
fi

if [[ $* == *--auto* ]]; then
  auto="y"
else
  auto="n"
fi

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

# echo ${cert} | jq
  echo ${cert} | jq -r '.data.certificate' > ../security/tls/$1.pem
  echo ${cert} | jq -r '.data.private_key' > ../security/tls/$1.key
}

set -e

setIfEmpty VAULT_ADDR "http://127.0.0.1:7777" "${VAULT_ADDR}"
setIfEmpty VAULT_TOKEN "myroot" "${VAULT_TOKEN}"

cd $(dirname $0)
mkdir -p ../security
mkdir -p ../security/ca
mkdir -p ../security/tls

echo "=> Root CA informations (will be used for TLS)"
createCA ca

echo "=> Users root CA informations (will be used for user certs)"
createCA users-ca

echo "=> Agents root CA informations (will be used for agent certs)"
createCA agents-ca

vault secrets enable pki
vault secrets tune -max-lease-ttl=53000h pki

caCert=$(cat ../security/ca/ca.pem | sed -e 's/$/\\n/g' | tr -d '\n')$'\\n'
caKey=$(cat ../security/ca/ca.key | sed -e 's/$/\\n/g' | tr -d '\n')

# Set root CA
curl -sSL \
  -H "X-Vault-Token: $VAULT_TOKEN" \
  -X POST \
  -d "{\"pem_bundle\": \"${caCert}${caKey}\"}" \
  "$VAULT_ADDR/v1/pki/config/ca"

# Create role
curl -sSL \
  -H "X-Vault-Token: $VAULT_TOKEN" \
  -X POST \
  -d "{\"allow_any_name\": true, \"max_ttl\": \"52560h\", \"organization\": \"$1\"}" \
  "$VAULT_ADDR/v1/pki/roles/$1"

createTLSCert banana.enix.io $1
createTLSCert api.banana.enix.io $1

cat ../security/ca/users-ca.pem > ../security/ca/trusted-ca.pem
echo >> ../security/ca/trusted-ca.pem
cat ../security/ca/agents-ca.pem >> ../security/ca/trusted-ca.pem

./init_vault.sh
./init_storage_access.sh

cd - > /dev/null

docker-compose restart proxy
echo $'\n=> ready'
