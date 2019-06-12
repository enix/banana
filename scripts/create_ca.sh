#! /usr/bin/env bash

openssl genrsa -des3 -out ./$1-tmp.key 2048
openssl req -x509 -new -nodes -extensions v3_ca -key ./$1-tmp.key -sha256 -days 2190 -out ./$1.pem
openssl rsa -in ./$1-tmp.key -out ./$1.key -outform pem
rm ./$1-tmp.key
