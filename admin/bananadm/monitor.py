from bananadm import vault
from bananadm import policies


def init(args):
    client = vault.get_vault_client(args)
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
    print(token)
