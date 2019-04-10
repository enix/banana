#! /usr/bin/env bash

basedir=$(pwd)
endpoint="$1"
filename="${basedir}/$2"
username="$3"

cd $(dirname $0)

if [[ $# -lt 3 ]] ; then
  echo "Usage: sign <endpoint> <data file> <username>"
  exit 1
fi

./sign.sh ${filename} ../security/out/${username}.key
signature="$(cat signature.sha256 | tr -d '\n')"
rm signature.sha256

curl \
  -X POST \
  -H "X-Signature: ${signature}" \
  -d "@${filename}" \
  --cacert ../security/ca/agents-ca.pem \
  --cert ../security/out/${username}.pem \
  --key ../security/out/${username}.key \
  "https://api.banana.enix.io:7443${endpoint}" \
  ${@:4}

cd - > /dev/null
