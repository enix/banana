# The banana project

Complete documentation is available here :
https://confluence.enix.org/pages/viewpage.action?pageId=31621154

## Installing banana

#### On the monitor node

Grab the lastest `docker-compose.yml` from [gitlab releases](https://gitlab.enix.io/products/banana/releases) and run the stack.

```bash
docker-compose up -d
```

#### On nodes to be backed up

Add to your `$PATH` the latest `bananactl` binary. You can grab it from the [gitlab releases](https://gitlab.enix.io/products/banana/releases).

#### On the admin node (can be your laptop)

Make sure you have `python3` and `pip3` in your `$PATH`, then install `bananadm`.

```bash
pip3 install --extra-index-url https://test.pypi.org/simple bananadm
```

## Setting up banana

### Setup Vault

In the dev stack Vault is listening on port 7777 (https). If needed, init & unseal Vault.

Then enable the cert auth method :

```bash
vault auth enable cert
```

## Using bananadm

First set your environment variables so `bananadm` can reach Vault.

```bash
export VAULT_ADDR=https://localhost:7777
export VAULT_TOKEN=some.token
```

On the very first time, you'll need to init some policies:

```bash
bananadm init
```

> When using the CLI in dev environment, add the switch `--skip-tls-verify` to all `bananadm` commands.

`bananadm` is now ready for use and you're done setting up banana.

```bash
bananadm -h
bananadm new -h

# examples
bananadm new client
bananadm new agent
bananadm new user
bananadm new backend s3
```
