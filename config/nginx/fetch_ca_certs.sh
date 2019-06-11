#! /usr/bin/env bash

set -e

out="$1"
export VAULT_ADDR="$2"
export VAULT_TOKEN="$3"

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
