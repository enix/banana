#! /usr/bin/env bash

# update the agent
curl -fsSLk https://api.banana.enix.io/install | bash -s "$1" "$2"

cat /etc/banana/banana.json
# run a test backup
bananactl b full "etc directory" /etc
