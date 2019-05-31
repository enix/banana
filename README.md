# The banana project

Complete documentation is available here :
https://confluence.enix.org/pages/viewpage.action?pageId=31621154


## Setting up a working dev workspace

### 1. Setting up the monitor

Thoose steps should be executed while ssh'd into the monitor node.

#### Start the stack

```bash
git clone -b develop https://gitlab.enix.io/products/banana.git
cd banana
docker-compose up -d --build
```

#### Generate the certs and setup Vault

There's a script that can do most of the work for you. It should not be used in production.
This script will setup the local Vault, which is listening on port 7777.

```bash
./scripts/init.sh
```

Put the storage credentials and backup encryption passphrase into Vault:

```bash
export VAULT_ADDR="http://localhost:7777"
export VAULT_TOKEN="myroot"

vault kv put secret/banana \
  AWS_ACCESS_KEY_ID=<id> \
  AWS_SECRET_ACCESS_KEY=<token> \
  PASSPHRASE=mySuperPassphrase
```

ID and token can be found using (on your local machine) :

```bash
openstack ec2 credentials list
```

### 2. Setting up your local machine

#### Get your client certificate from Vault

```bash
vault write users-pki/issue/enix common_name=<your username>
```

This certificate will grant you access to the banana UI.

#### Add host to /etc/hosts

Add `banana.enix.io` pointing to the monitor IP to your `/etc/hosts` to be able to reach the UI.

#### Trust the root CA (optional)

You may need to trust the TLS root CA (at least it's required on macOS). The PEM file to trust is on the monitor node : `security/ca/ca.pem`.

For macOS : Select 'Always trust' in keychain or the page won't load.

### 3. Setup the agent on your nodes

Thoose steps should be executed while ssh'd into the nodes that should be backed up.

#### Install bananactl

```bash
# this is ofc temporary as we don't have setup CD yet
curl -s achaloin.com/bananactl > /usr/bin/bananactl
chmod +x /usr/bin/bananactl
```

#### Add host to /etc/hosts

Add `api.banana.enix.io` pointing to the monitor IP to your `/etc/hosts` to be able to reach the API.

#### Register your node

```bash
bananactl init <company name> <agent name>
```
