#! /usr/bin/env sh

subject="/C=FR/ST=Ile-De-France/L=Paris/O=Enix/OU=Banana/CN=banana.dev.enix.io"
cert="/tls/banana.dev.enix.io.pem"
key="/tls/banana.dev.enix.io.key"

mkdir -p /data && chown vault:vault /data
mkdir -p /tls && chown vault:vault /tls

apk add --update openssl

openssl req \
	-new \
	-newkey rsa:4096 \
	-days 365 \
	-nodes \
	-x509 \
	-subj ${subject} \
	-keyout ${key} \
	-out ${cert}

vault server -config /settings/vault/vault.hcl
