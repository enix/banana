import os
from libbananadm import vault


def create_agent(args):
    print('generating temporary token to allow new agent(s) to register')
    client = vault.get_vault_client(args)
    token = client.create_token(
        policies=['{}-{}-agent-creation'.format(args.root_path, args.client)],
        lease='1h',
    )

    print('success! join your new agent(s) using:\n')
    print(
        'bananactl {} --vault-addr={} init {} {} <agent name>'
        .format(
            '--skip-tls-verify' if args.skip_tls_verify else '',
            os.getenv('VAULT_ADDR'),
            token['auth']['client_token'],
            args.client,
        )
    )
