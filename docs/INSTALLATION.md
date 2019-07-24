# Installing banana

The following assumes that this domain is pointing to the monitor IP:

* `banana.dev.enix.io`

If this is not the case, you should edit your `/etc/hosts` file in consequence :

```
<ip of the monitor node>	banana.dev.enix.io
```

#### On the monitor node

Grab the lastest `docker-compose.yml` from [gitlab releases](https://gitlab.enix.io/products/banana/releases) and run the stack.

```bash
docker-compose up -d
```

#### On the admin node (can be your laptop)

Make sure you have `python3` and `pip3` installed, then install `bananadm`.

```bash
pip3 install bananadm
```

## Setting up Vault

You need to allow `bananadm` to interact with Vault. To do so :

> DISCLAIMER: For now, it is highly recommended to run this project on a dedicated Vault. The permissions granted to `bananadm` are dangerous and can lead to a full privilege escalation on your Vault instance.

If your Vault is already up and running, you can skip to step 3.

1. Init Vault :

```bash
export VAULT_ADDR=https://banana.dev.enix.io:7777

# for the sake of simplicity we use a single unseal key. for production, it is highly recommended to use more
vault operator init -tls-skip-verify -key-shares=1 -key-threshold=1
```

2. Unseal Vault :

```bash
# will prompt for the unseal key, which is the base 64 key in the previous command's output
vault operator unseal -tls-skip-verify
```

3. Log into Vault using any method. One possibility is to set the `VAULT_TOKEN` environment variables, just like this :

```bash
export VAULT_TOKEN=s.the_root_token_in_step_1_output
```

4. Download [the bananadm policy](https://gitlab.enix.io/products/banana/raw/master/config/vault/bananadm-policy.hcl).

5. Write this policy into Vault :

```bash
vault policy write -tls-skip-verify bananadm bananadm-policy.hcl
```

6. Issue a token with the associated permissions :

```bash
vault token create -tls-skip-verify -policy=bananadm
```

7. Downgrade your privileges by updating your Vault token with the newly generated token :

```bash
export VAULT_TOKEN=s.freshly_generated_bananadm_token
```

## Using bananadm

Make sure `VAULT_ADDR` and `VAULT_TOKEN` environment variables are set.

When using the CLI in dev environment, add the switch `--tls-skip-verify` to all `bananadm` commands :

```bash
alias bananadm="bananadm --tls-skip-verify"
```

On the very first time, you'll need to init some stuff:

```bash
bananadm init
```

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

* Install `bananactl` on the node(s):

First, install `curl`. Then install `bananactl` using the following command (add `-k` tu `curl` command if needed):

```bash
$ curl -fsS https://banana.dev.enix.io/install | bash -s - '<gitlab access token>'

downloading latest agent release...
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 5028k  100 5028k    0     0  25.3M      0 --:--:-- --:--:-- --:--:-- 25.3M
[...]
success!
```

* Create an agent join command using `bananadm`:

```bash
$ bananadm new agent

client in which create the agent(s)? enix
generating temporary token to allow new agent(s) to register
success! join your new agent(s) using:

bananactl --vault-addr=https://banana.dev.enix.io:7777 init s.BVYt1Hj3eLn6NDPS2fJIKzfO enix <agent name>
```

> The `new agent` command generates a token with a 1h TTL which has the required permissions to issue certificates from the client's agents PKI. The generated `bananactl` command can then be runned an unlimited amount of times to register multiple agents, while the token is still valid.

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
$ bananadm new user

client in which create the user? enix
username? arthur
issuing certificate with CN 'arthur' using PKI banana/enix/users-pki
creating p12 file
Enter Export Password: ****
Verifying - Enter Export Password: ****
successfully wrote arthur.p12
```

* Open your bowser on [banana.dev.enix.io](https://banana.dev.enix.io) and authenticate using the generated p12 file.
