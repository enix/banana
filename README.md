# The banana project

Complete documentation is available here :
https://confluence.enix.org/pages/viewpage.action?pageId=31621154

## Installing banana

The following assumes that thoose domains are pointing to the monitor IP:

* `banana.enix.io`
* `api.banana.enix.io`
* `vault.banana.enix.io`

If this is not the case, you should edit your `/etc/hosts` file in consequence :

```
<ip of the monitor node>	banana.enix.io api.banana.enix.io vault.banana.enix.io
```

#### On the monitor node

Grab the lastest `docker-compose.yml` from [gitlab releases](https://gitlab.enix.io/products/banana/releases) and run the stack.

```bash
docker-compose up -d
```

#### On the admin node (can be your laptop)

Make sure you have `python3` and `pip3` installed, then install `bananadm`.

```bash
pip3 install --extra-index-url https://test.pypi.org/simple bananadm
```

## Setting up Vault

If needed, init & unseal Vault. In the dev stack, Vault is listening on port 7777 (https).

## Using bananadm

First set your environment variables so `bananadm` can reach Vault.

```bash
export VAULT_ADDR=https://vault.banana.enix.io:7777
export VAULT_TOKEN=some.token
```

> When using the CLI in dev environment, add the switch `--skip-tls-verify` to all `bananadm` commands.

On the very first time, you'll need to init some stuff:

```bash
bananadm init
```

`bananadm` is now ready for use and you're done setting up banana.

```bash
bananadm -h
bananadm new -h
```

## Example: Setting up a node for backup

### Setup a client and storage backend

* Create a client in which the agent(s) will be registered:

```bash
$ bananadm new client

client name? enix
mounted KV at path 'banana/enix/secrets'
mounted root PKI at path 'banana/enix/root-pki'
mounted users intermediate PKI at path 'banana/enix/users-pki'
mounted agents intermediate PKI at path 'banana/enix/agents-pki'
created policy enix-agent-creation
created policy enix-agent-access
allowed agent certs to login into vault
successfully reloaded nginx trust configuration
successfully created client 'enix'
```

* Set some storage credentials:

```bash
$ bananadm new backend s3

client in which create the secret? enix
backend name? openstack
AWS_ACCESS_KEY_ID? 68c8e041************************
AWS_SECRET_ACCESS_KEY? 28664564************************
successfully saved secret backends/openstack in KV engine banana/enix/secrets
```

All agents in the same client share all the storage secrets.

### Setup the agent(s)

* Install `bananactl` on the node(s):

First, install `curl`. Then install `bananactl` using the following command (add `-k` tu `curl` command if needed):

```bash
$ curl -fsS https://api.banana.enix.io/install | bash -s - '<gitlab access token>'

downloading latest agent release...
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 5028k  100 5028k    0     0  25.3M      0 --:--:-- --:--:-- --:--:-- 25.3M
$ unzip -o agent.zip
Archive:  agent.zip
  inflating: bananactl-linux
  inflating: config/systemd/banana.service
  inflating: config/systemd/banana.timer
$ cp bananactl-linux /usr/bin/bananactl
$ cp config/systemd/banana.service config/systemd/banana.timer /etc/systemd/system/
$ systemctl start banana.timer
$ systemctl enable banana.timer
success!
```

* Create an agent join command using `bananadm`:

```bash
$ bananadm new agent

client in which create the agent(s)? enix
generating temporary token to allow new agent(s) to register
success! join your new agent(s) using:

bananactl --vault-addr=https://vault.banana.enix.io:7777 init s.BVYt1Hj3eLn6NDPS2fJIKzfO enix <agent name>
```

* Copy/paste the command on each node that you'd like to initialize.

* Edit the `/etc/banana/schedule.json` file on each node with your backup configuration:

```json
{
	"my-backup-name": {
		"target": "/etc",
		"interval": 0.042,
		"full_every": 2
	}
}
```

This example configuration will :

* backup the directory `/etc` as `my-backup-name`
* run every hour (1 / 24 ~= 0.042)
* do a full backup every 2 backups (so half of the backups will be incrementals)

## Example: Accessing the banana UI

* Create a user :

```bash
$ bananadm new user

client in which create the user? enix
username? arthur
issuing certificate with CN 'arthur' using PKI banana/enix/users-pki
creating p12 file
Enter Export Password: ****
Verifying - Enter Export Password: ****
successfully wrote arthur.p12
```

* Open your bowser on [banana.enix.io](https://banana.enix.io) and authenticate using the generated p12 file.
