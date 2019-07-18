#! /usr/bin/env bash

curl -fsS https://api.banana.enix.io/install | bash -s "$1"
bananactl version
