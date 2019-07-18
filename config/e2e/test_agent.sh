#! /usr/bin/env bash

echo $1
curl -fsSLk https://api.banana.enix.io/install | bash -s "$1"
bananactl version
