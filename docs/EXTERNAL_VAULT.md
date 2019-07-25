## Manually setup a Vault instance for Banana

If your Vault is already setup, you can skip to step 3.

1. Init Vault :

```bash
export VAULT_ADDR=https://banana.dev.enix.io:8200

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

4. Download [the bananadm policy](../config/vault/bananadm-policy.hcl).

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
