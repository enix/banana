#! /usr/bin/env bash

set -e
source ~/.ssh/environment
sudo -E bash
cd ~

# upgrade the stack
docker-compose -f /root/docker-compose.yml up -d

# unseal vault
vault operator unseal -tls-skip-verify ${VAULT_UNSEAL_KEY}

# reload nginx configuration
bananadm --tls-skip-verify reconfigure
