# Installing banana

The following assumes that this domain is pointing to the monitor IP:

* `banana.dev.enix.io`

If this is not the case, you should edit your `/etc/hosts` file in consequence :

```
<ip of the monitor node>	banana.dev.enix.io
```

### On the monitor node

Grab the lastest `docker-compose.yml` from [github releases](https://github.com/enix/banana/releases) and run the stack.

```bash
docker-compose up -d
```

### On the admin node (can be your laptop)

* Using docker

```bash
docker run --rm -it enix/bananadm -h
```

* Using pip

```bash
pip3 install bananadm
```

## Setting up Vault

You need to allow `bananadm` to interact with Vault.

> DISCLAIMER: For now, it is highly recommended to run this project on a dedicated Vault. The permissions granted to `bananadm` are dangerous and can lead to a full privilege escalation on your Vault instance.

If you'd like to setup banana with an existing or external Vault instance, please follow [this guide](SETUP_VAULT.md). Otherwise, if you prefer using the builtin Vault, you can init everything with a single command (see below).

## Using bananadm

When using the CLI in dev environment, add the switch `--tls-skip-verify` to all `bananadm` commands :

```bash
alias bananadm="bananadm --tls-skip-verify"
```

On the very first time, you'll need to init some stuff.

#### Using the builtin Vault

This command will do everything needed to get Vault running, including initialization, unseal and required policies upload.

```
export VAULT_ADDR=banana.dev.enix.io:8200
bananadm init --from-scratch
```

#### Using an external Vault

Make sure `VAULT_ADRR` and `VAULT_TOKEN` environment variables are set.

```bash
bananadm init
```

---

`bananadm` is now ready for use and you're done setting up banana.

```bash
bananadm -h
bananadm new -h
bananadm new backend -h
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
created policy banana-enix-agent-creation
created policy banana-enix-agent-access
allowed agent certs to login into vault
successfully reloaded nginx trust configuration
successfully created client 'enix'
```

* Set some storage credentials:

> By default, the agent will try to use the backend named 'openstack'. If you'd like to use a different name, you should edit in consequence the `vault.storage_secret_path` key in `banana.json` _for each agent_.

```bash
$ bananadm new backend s3

client in which create the secret? enix
backend name? openstack
AWS_ACCESS_KEY_ID? 68c8e041************************
AWS_SECRET_ACCESS_KEY? 28664564************************
successfully saved secret backends/openstack in KV engine banana/enix/secrets
```

All agents in the same client share all the storage secrets. By default (for now), agents will try to use credentials stored in the backend named 'openstack'.

### Setup the agent(s)

* Install `bananagent` on the node(s):

```
echo 'deb [trusted=yes] https://raw.githubusercontent.com/enix/packages/master' unstable main >> /etc/apt/sources.list
apt update
apt install bananagent
```

* Create an agent join command using `bananadm`:

```bash
$ bananadm new agent

client in which create the agent(s)? enix
generating temporary token to allow new agent(s) to register
success! join your new agent(s) using:

bananagent --vault-addr=https://banana.dev.enix.io:8200 init s.BVYt1Hj3eLn6NDPS2fJIKzfO enix <agent name>
```

> The `new agent` command generates a token with a 1h TTL which has the required permissions to issue certificates from the client's agents PKI. The generated `bananagent` command can then be runned an unlimited amount of times to register multiple agents, while the token is still valid.

* Copy/paste the command on each node that you'd like to initialize.

* Edit the `/etc/banana/schedule.json` file on each node with your backup configuration:

```json
{
	"my-backup-name": {
		"interval": 0.042,
		"full_every": 2,
		"plugin_args": [
			"/etc"
		]
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
$ bananadm new user > cert.p12

client in which create the user? enix
username? arthur
issuing certificate with CN 'arthur' using PKI banana/enix/users-pki
creating p12 file
Enter Export Password: ****
Verifying - Enter Export Password: ****
successfully wrote p12 data to stdout
```

* Open your bowser on [banana.dev.enix.io](https://banana.dev.enix.io) and authenticate using the generated p12 file.
