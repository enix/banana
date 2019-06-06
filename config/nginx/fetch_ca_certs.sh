#! /usr/bin/env sh

set -e

out="$1"
VAULT_TOKEN="$2"

vault token lookup > /dev/null

rm -f ${out}
mounts=$(vault read /sys/mounts -format=json | jq -r '.data | with_entries(select(.key | match("banana.*pki"))) | keys[]')
for mount in ${mounts}; do
	cacert=$(vault read ${mount}/cert/ca -format=json | jq -r '.data.certificate')
	echo "${cacert}" >> ${out}
done

echo 'wrote CA certs to' ${out}
