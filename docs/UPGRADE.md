# Upgrading banana to the latest version

## Update the monitor stack

* Grab the lastest `docker-compose.yml` from [gitlab releases](https://gitlab.enix.io/products/banana/releases).
* Override the old compose file with the newly downloaded one.
* Run `docker-compose up -d`

Please note that :

* Vault's TLS certificate will be re-generated
* You'll need to unseal Vault again by providing the master key(s)

## Update the agent binary

Run the install script again.

```bash
curl -fsS https://api.banana.enix.io/install | bash -s - '<gitlab access token>'
```
