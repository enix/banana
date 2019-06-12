#! /usr/bin/env bash

out="$1"

read vault_addr
read vault_token
export VAULT_ADDR="${vault_addr}"
export VAULT_TOKEN="${vault_token}"

vault token lookup -tls-skip-verify > /dev/null 2>&1
if [ $? -ne 0 ]; then
	export VAULT_ADDR="https://vault:8200"
fi

set -e

vault token lookup -tls-skip-verify > /dev/null

mounts=$(vault read -tls-skip-verify /sys/mounts -format=json | jq -r '.data | with_entries(select(.key | match("banana.*pki"))) | keys[]')
if [ -z "${mounts}" ]; then
	exit
fi

cp /dev/null ${out}
for mount in ${mounts}; do
	cacert=$(vault read -tls-skip-verify ${mount}/cert/ca -format=json | jq -r '.data.certificate')
	echo "${cacert}" >> ${out}
done

echo 'wrote CA certs to' ${out}
sudo /usr/local/openresty/nginx/sbin/nginx -s reload
