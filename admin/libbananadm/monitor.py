import os
import requests
from libbananadm import vault
from libbananadm import policies
from hvac import exceptions


def init(args):
    if args.from_scratch:
        root_token, keys, token = init_from_scratch(args)

    client = vault.get_vault_client(args)

    try:
        client.sys.enable_auth_method(
            method_type='cert',
            path='{}/cert'.format(args.root_path),
        )
        print(
            'enabled cert auth method on path /auth/{}/cert'
            .format(args.root_path)
        )
    except exceptions.InvalidRequest as error:
        if 'path is already in use' in error.errors[0]:
            print(
                'error: looks like this vault was already setup for banana, '
                + 'aborting'
            )
            exit(1)
        raise error

    monitor_policy_name = 'banana-monitor'
    client.sys.create_or_update_policy(
        name=monitor_policy_name,
        policy=policies.generate_monitor_policy(args),
    )
    print('created policy {}'.format(monitor_policy_name))

    if args.from_scratch:
        print(
            '\n==========================\n'
            'Your Vault instance is now ready to go.\n'
            'To be able to reach it, '
            'you should add this to your shell rc file :\n\n'
            'export VAULT_ADDR={}\n'
            'export VAULT_TOKEN={}\n\n'
            'Here is your Vault unseal keys, you should store them safely:\n\n'
            '{}\n\n'
            'Your Vault root token is: {} (but you shouldn\'t need it)'
            .format(
                os.getenv('VAULT_ADDR'),
                token,
                '\n'.join(keys),
                root_token,
            )
        )


def init_from_scratch(args):
    client = vault.get_vault_client(args)

    # TODO: Use more shares/threshold for security
    res = client.sys.initialize(1, 1)
    root_token = res['root_token']
    keys = res['keys']
    print('initialized vault')

    client.sys.submit_unseal_keys(keys)
    print('unsealed vault')

    os.environ['VAULT_TOKEN'] = root_token
    client = vault.get_vault_client(args)

    res = requests.get(
        'https://raw.githubusercontent.com/enix/banana/master/'
        'config/vault/bananadm-policy.hcl',
    )
    if res.status_code >= 400:
        print(res.json())
        return
    policy = res.text
    print('downloaded bananadm policy from github')

    bananadm_policy_name = 'bananadm'
    client.sys.create_or_update_policy(
        name=bananadm_policy_name,
        policy=policy,
    )
    print('uploaded policy {}'.format(bananadm_policy_name))

    token_res = client.create_token(policies=['bananadm'])
    token = token_res['auth']['client_token']

    os.environ['VAULT_TOKEN'] = token
    return root_token, keys, token


def reconfigure(args):
    client = vault.get_vault_client(args)
    token = client.create_token(
        policies=['banana-monitor'],
        lease='1h',
    )

    res = requests.post(
        args.monitor_addr + '/reconfigure',
        data='{}\n{}'.format(
            os.getenv('VAULT_ADDR'),
            token['auth']['client_token'],
        ),
        verify=not args.tls_skip_verify,
    )
    if res.status_code >= 400:
        print(res.json())

    print('successfully reloaded nginx trust configuration')
