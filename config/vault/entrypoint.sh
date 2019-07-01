#! /usr/bin/env sh

subject="/C=FR/ST=Ile-De-France/L=Paris/O=Enix/OU=Banana/CN=vault.banana.enix.io"
cert="/tls/vault.banana.enix.io.pem"
key="/tls/vault.banana.enix.io.key"

openssl req \
	-new \
	-newkey rsa:4096 \
	-days 365 \
	-nodes \
	-x509 \
	-subj ${subject} \
	-keyout ${key} \
	-out ${cert}

vault server -config /vault.hcl
