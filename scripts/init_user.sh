#! /usr/bin/env bash

if [[ $# -lt 2 ]] ; then
  echo "Usage: $0 <org> <username>"
  exit 1
fi

function setIfEmpty {
  if [ -z "$3" ]; then
    export "$1"="$2"
  fi
}

cd $(dirname $0)

setIfEmpty VAULT_ADDR "http://127.0.0.1:7777" "${VAULT_ADDR}"
setIfEmpty VAULT_TOKEN "myroot" "${VAULT_TOKEN}"

mkdir -p ../security/out

# Create role
curl -sSL \
  -H "X-Vault-Token: $VAULT_TOKEN" \
  -X POST \
  -d "{\"allow_any_name\": true, \"max_ttl\": \"52560h\", \"organization\": \"$1\"}" \
  "$VAULT_ADDR/v1/users-pki/roles/$1"

# Issue cert
cert=$(curl -sSL \
  -H "X-Vault-Token: $VAULT_TOKEN" \
  -X POST \
  -d "{\"common_name\": \"$2\", \"ttl\": \"52500h\"}" \
  "$VAULT_ADDR/v1/users-pki/issue/$1")

echo ${cert} | jq -r '.data.certificate' > ../security/out/$2.pem
echo ${cert} | jq -r '.data.private_key' > ../security/out/$2.key

openssl pkcs12 -export -out ../security/out/$2.p12 -inkey ../security/out/$2.key -in ../security/out/$2.pem

echo "saved ../security/out/$2.p12"

cd -
