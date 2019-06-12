FROM openresty/openresty:1.15.8.1-1-stretch

RUN apt update

RUN apt install -y zip jq sudo

ADD https://releases.hashicorp.com/vault/1.1.3/vault_1.1.3_linux_amd64.zip /vault.zip

RUN unzip /vault.zip

RUN chmod +x /vault && mv /vault /usr/bin/vault

RUN echo 'nobody ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers

RUN openssl genrsa -des3 -out tmp.key -passout pass:enix 2048 \
 && echo 'FR\nIle-De-France\nParis\nEnix\n.\nenix.io\nenix@enix.io' \
  | openssl req -x509 -new -nodes -extensions v3_ca -key tmp.key -sha256 -days 2190 -passin pass:enix -out trusted-ca.pem \
 && rm tmp.key \
 && chmod 666 /trusted-ca.pem

COPY ./tls /tls

COPY ./nginx/conf.d /etc/nginx/conf.d

COPY ./nginx/fetch_ca_certs.sh /fetch_ca_certs.sh
