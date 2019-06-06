import os
from bananadm import vault


def create_agent(args):
    client = vault.get_vault_client(args)
    token = client.create_token(
        policies=['{}-agent-creation'.format(args.client)],
        lease='1h',
    )

    print('Join your new agent(s) using:\n')
    print(
        'bananactl {} --vault-addr={} init {} {} <agent name>'
        .format(
            '--skip-tls-verify' if args.skip_tls_verify else '',
            os.getenv('VAULT_ADDR'),
            token['auth']['client_token'],
            args.client,
        )
    )
