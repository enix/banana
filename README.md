# The banana project

Complete documentation is available here :
https://confluence.enix.org/pages/viewpage.action?pageId=31621154


## Setting up banana

### Start the stack

```bash
git clone -b develop https://gitlab.enix.io/products/banana.git
cd banana
docker-compose up -d --build
```

### Setup Vault

In the dev stack Vault is listening on port 7777 (https). If needed, init & unseal Vault.

Then enable the cert auth method :

```bash
vault auth enable cert
```

### Add core policies needed for banana

The switch `--skip-tls-verify` will be needed in dev environment.

```bash
bananadm init
```

## Using bananadm

```bash
bananadm -h
bananadm new -h

# examples
bananadm new client
bananadm new agent
bananadm new user
bananadm new backend s3
```
