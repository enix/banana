#! /usr/bin/env bash

set -e

echo $2
# update the agent
# curl -fsSLk https://api.banana.enix.io/install | bash -s "$1" "$2"
cp /home/ubuntu/bananactl-linux /usr/local/bin/bananactl
echo version that will be tested: $(bananactl version)

cat /etc/banana/banana.json
# run a test backup
env | grep VAULT
bananactl b full "etc directory" /etc
