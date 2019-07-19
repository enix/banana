#! /usr/bin/env bash

# update the agent
curl -fsSLk https://api.banana.enix.io/install | bash -s "$1"

# run a test backup
bananactl b full "etc directory" /etc
