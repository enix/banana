#! /usr/bin/env bash

filename=$1
privatekey=$2

if [[ $# -lt 2 ]] ; then
  echo "Usage: sign <file> <private_key>"
  exit 1
fi

openssl dgst -sha256 -sign $privatekey -out $filename.sha256 $filename
openssl base64 -in $filename.sha256 -out signature.sha256

rm $filename.sha256
