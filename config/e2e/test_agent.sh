#! /usr/bin/env bash

set -e

# update the agent
cp /home/ubuntu/bananagent-linux /usr/local/bin/bananagent
echo version that will be tested: $(bananagent version)

# run a test backup
bananagent b full "etc directory" /etc
