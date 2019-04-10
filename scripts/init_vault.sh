#! /usr/bin/env bash

function setIfEmpty {
  if [ -z "$3" ]; then
    export "$1"="$2"
  fi
}

cd $(dirname $0)

setIfEmpty VAULT_ADDR "http://127.0.0.1:7777" "${VAULT_ADDR}"
setIfEmpty VAULT_TOKEN "myroot" "${VAULT_TOKEN}"

vault secrets enable --path=users-pki pki
vault secrets tune -max-lease-ttl=53000h users-pki

usersCaCert=$(cat ../security/ca/users-ca.pem | sed -e 's/$/\\n/g' | tr -d '\n')$'\\n'
usersCaKey=$(cat ../security/ca/users-ca.key | sed -e 's/$/\\n/g' | tr -d '\n')

# Set root CA
curl -sSL \
  -H "X-Vault-Token: $VAULT_TOKEN" \
  -X POST \
  -d "{\"pem_bundle\": \"${usersCaCert}${usersCaKey}\"}" \
  "$VAULT_ADDR/v1/users-pki/config/ca"

vault secrets enable --path=agents-pki pki
vault secrets tune -max-lease-ttl=53000h agents-pki

agentsCaCert=$(cat ../security/ca/agents-ca.pem | sed -e 's/$/\\n/g' | tr -d '\n')$'\\n'
agentsCaKey=$(cat ../security/ca/agents-ca.key | sed -e 's/$/\\n/g' | tr -d '\n')

# Set root CA
curl -sSL \
  -H "X-Vault-Token: $VAULT_TOKEN" \
  -X POST \
  -d "{\"pem_bundle\": \"${agentsCaCert}${agentsCaKey}\"}" \
  "$VAULT_ADDR/v1/agents-pki/config/ca"

cd -  > /dev/null
