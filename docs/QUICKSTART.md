## Quickstart guide

1. Run the monitor stack using docker-compose. The compose file can be found on [github releases](https://github.com/enix/banana/releases).

2. Setup your banana instance:

```bash
# initialize vault and get your bananadm token
docker run --rm -ie VAULT_ADDR=https://banana.dev.enix.io:8200 enix/bananadm:1.12.5 --tls-skip-verify init --from-scratch

# for the sake of simplicity
alias bananadm="docker run --rm -ie VAULT_ADDR=https://banana.dev.enix.io:8200 -e VAULT_TOKEN=<s.TOKEN_IN_PREV_CMD> enix/bananadm:1.12.5 --tls-skip-verify"

# create a client
bananadm new client

# register storage backend credentials
bananadm new backend s3

# create a certifiate for your user to access the banana UI
bananadm new user > cert.pem

# print a command to be runned on nodes to register them
bananadm new agent
```

3. Install bananagent on your node(s).

```bash
echo 'deb [trusted=yes] https://raw.githubusercontent.com/enix/packages/master' unstable main >> /etc/apt/sources.list
apt update
apt install bananagent
```

4. Copy/paste the init command from step 2.

5. Go to [the banana UI](https://banana.dev.enix.io/scheduler), generate a backup schedule, and write it to `/etc/banana/schedue.json`.

You're ready to go!
