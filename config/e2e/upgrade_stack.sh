#! /usr/bin/env bash

set -e
sudo -i

# upgrade the stack
docker-compose -f /root/docker-compose.yml up -d

# unseal vault
env | grep VAULT
echo vault operator unseal ${VAULT_UNSEAL_KEY} -tls-skip-verify
vault operator unseal ${VAULT_UNSEAL_KEY} -tls-skip-verify

# reload nginx configuration
bananadm --tls-skip-verify reconfigure
