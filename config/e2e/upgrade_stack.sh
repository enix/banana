#! /usr/bin/env bash

set -e
sudo -i
source ~/.ssh/environment

# upgrade the stack
docker-compose -f /root/docker-compose.yml up -d

# unseal vault
vault operator unseal ${VAULT_UNSEAL_KEY} -tls-skip-verify

# reload nginx configuration
bananadm --tls-skip-verify reconfigure