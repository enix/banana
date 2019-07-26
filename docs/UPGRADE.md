# Upgrading banana to the latest version

## Update the monitor stack

1. Grab the lastest `docker-compose.yml` from [github releases](https://github.com/enix/banana/releases).
2. Override the old compose file with the newly downloaded one.
3. Run `docker-compose up -d`.
4. If needed, unseal Vault again by providing the master key(s) : `vault operator unseal -tls-skip-verify`
5. If you installed it with pip, upgrade bananadm: `pip install bananadm -U`
6. Run `bananadm --tls-skip-verify reconfigure`

Please note that Vault's TLS certificate will be re-generated.

## Update the agent binary

You can update using apt:

```
apt update
apt install bananagent
```
