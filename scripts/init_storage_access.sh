#! /usr/bin/env bash

function setIfEmpty {
  if [ -z "$3" ]; then
    export "$1"="$2"
  fi
}

echo "taking storage credentials from openstack, make sure you have at least one generated"

setIfEmpty VAULT_ADDR "http://127.0.0.1:7777" "${VAULT_ADDR}"
setIfEmpty VAULT_TOKEN "myroot" "${VAULT_TOKEN}"

access=$(openstack --os-interface public ec2 credentials list --format value | head -n 1)
id=$(echo "${access}" | awk '{ print $1 }')
token=$(echo "${access}" | awk '{ print $2 }')

echo "your access token is:" ${id}
echo "your secret token is:" ${token}

vault kv put secret/storage_access \
  AWS_ACCESS_KEY_ID=${id} \
  AWS_SECRET_ACCESS_KEY=${token} \
  PASSPHRASE=mySuperPassphrase
