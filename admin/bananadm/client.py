import os
import requests
from bananadm import vault
from bananadm import policies


def create_client(args):
    client = vault.get_vault_client(args)
    client_pki = '{}/{}/root-pki'.format(args.root_path, args.name)
    client_kv = '{}/{}/secrets'.format(args.root_path, args.name)
    client_users_pki = '{}/{}/users-pki'.format(args.root_path, args.name)
    client_agents_pki = '{}/{}/agents-pki'.format(args.root_path, args.name)
    client.sys.enable_secrets_engine('kv', path=client_kv)

    client.sys.enable_secrets_engine('pki', path=client_pki, config={
        'max_lease_ttl': '43800h',
    })
    client.secrets.pki.generate_root(
        type='internal',
        common_name=args.name,
        mount_point=client_pki,
        extra_params={
            'ttl': '43800h',
        },
    )

    client.create_intermediate_ca(
        args, client_pki, client_users_pki, 'User',
    )
    agents_cert = client.create_intermediate_ca(
        args, client_pki, client_agents_pki, 'Agent',
    )

    client.sys.create_or_update_policy(
        name='{}-agent-creation'.format(args.name),
        policy=policies.generate_agent_install_policy(args),
    )
    client.sys.create_or_update_policy(
        name='{}-agent-access'.format(args.name),
        policy=policies.generate_agent_access_policy(args),
    )

    res = requests.post(
        '{}/v1/auth/cert/certs/{}'.format(
            os.getenv('VAULT_ADDR'),
            args.name,
        ),
        json={
            'display_name': args.name,
            'policies': '{}-agent-access'.format(args.name),
            'certificate': agents_cert,
        },
        headers={
            'X-Vault-Token': os.getenv('VAULT_TOKEN'),
        },
        verify=not args.skip_tls_verify,
    )
    if res.status_code >= 400:
        print(res.json())
        print('did you enable cert auth method in vault?')
        print('\n$ vault auth enable cert')
        exit(1)
