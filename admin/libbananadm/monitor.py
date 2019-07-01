import os
import requests
from libbananadm import vault
from libbananadm import policies


def init(args):
    client = vault.get_vault_client(args)
    monitor_policy_name = 'banana-monitor'
    client.sys.create_or_update_policy(
        name=monitor_policy_name,
        policy=policies.generate_monitor_policy(args),
    )
    print('created policy {}'.format(monitor_policy_name))

    client.sys.enable_auth_method(
        method_type='cert',
        path='{}/cert'.format(args.root_path),
    )
    print(
        'enabled cert auth method on path /auth/{}/cert'.format(args.root_path)
    )


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
        verify=not args.skip_tls_verify,
    )
    if res.status_code >= 400:
        print(res.json())

    print('successfully reloaded nginx trust configuration')
