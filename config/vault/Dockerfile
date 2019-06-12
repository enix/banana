FROM vault

COPY ./vault /config

COPY ./tls /tls

CMD mkdir -p /data && chown vault:vault /data && vault server -config /config/vault.hcl
