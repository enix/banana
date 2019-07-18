#! /usr/bin/env bash

set -e
sudo -i

# upgrade the stack
docker-compose -f /root/docker-compose.yml up -d

# unseal vault
echo ${VAULT_UNSEAL_KEY} | vault operator unseal -tls-skip-verify

# reload nginx configuration
bananadm --tls-skip-verify reconfigure
