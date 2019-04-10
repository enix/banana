#! /usr/bin/env bash

filename=$1
signature=$2
cert=$3

if [[ $# -lt 3 ]] ; then
  echo "Usage: verify <file> <signature> <cert>"
  exit 1
fi

openssl x509 -pubkey -in ${cert} -noout > /tmp/${filename}.key
openssl base64 -d -in ${signature} -out /tmp/${filename}.sha256
openssl dgst -sha256 -verify /tmp/${filename}.key -signature /tmp/${filename}.sha256 ${filename}

rm /tmp/${filename}.sha256 /tmp/${filename}.key
