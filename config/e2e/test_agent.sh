#! /usr/bin/env bash

set -e

# update the agent
cp /home/ubuntu/bananactl-linux /usr/local/bin/bananactl
echo version that will be tested: $(bananactl version)

# run a test backup
bananactl b full "etc directory" /etc
