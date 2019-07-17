import os
import requests
from libbananadm import vault
from libbananadm import policies
from hvac import exceptions


def init(args):
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
