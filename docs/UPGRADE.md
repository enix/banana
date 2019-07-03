# Upgrading banana to the latest version

## Update the monitor stack

1. Grab the lastest `docker-compose.yml` from [gitlab releases](https://gitlab.enix.io/products/banana/releases).
2. Override the old compose file with the newly downloaded one.
3. Run `docker-compose up -d`
4. Unseal Vault again by providing the master key(s)
5. Run `bananadm reconfigure`

Please note that Vault's TLS certificate will be re-generated.

## Update the agent binary

Run the install script again.

```bash
curl -fsS https://api.banana.enix.io/install | bash -s - '<gitlab access token>'
```
