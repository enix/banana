#! /usr/bin/env bash

sudo -i

curl -fsSLk https://api.banana.enix.io/install | bash -s "$1"
bananactl version
