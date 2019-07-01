FROM vault

COPY ./vault.hcl /vault.hcl

COPY ./entrypoint.sh /entrypoint.sh

RUN mkdir -p /tls && chown vault:vault /tls

RUN mkdir -p /data && chown vault:vault /data

RUN apk add --update openssl

CMD /entrypoint.sh
