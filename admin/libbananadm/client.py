import os
import requests
from libbananadm import vault
from libbananadm import policies
from libbananadm import monitor
from tabulate import tabulate
from hvac import exceptions


def create_client(args):
    client = vault.get_vault_client(args)
    client_pki = '{}/{}/root-pki'.format(args.root_path, args.name)
    client_kv = '{}/{}/secrets'.format(args.root_path, args.name)
    client_users_pki = '{}/{}/users-pki'.format(args.root_path, args.name)
    client_agents_pki = '{}/{}/agents-pki'.format(args.root_path, args.name)

    try:
        client.sys.enable_secrets_engine('kv-v2', path=client_kv)
        print('mounted KV at path \'{}\''.format(client_kv))
    except exceptions.InvalidRequest as error:
        if 'existing mount' in error.errors[0]:
            print('error: client {} already exists'.format(args.name))
            exit(1)
        raise error

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
    print('mounted root PKI at path \'{}\''.format(client_pki))

    vault.create_intermediate_ca(
        args, client_pki, client_users_pki, 'User',
    )
    print(
        'mounted users intermediate PKI at path \'{}\''
        .format(client_users_pki)
    )

    agents_cert = vault.create_intermediate_ca(
        args, client_pki, client_agents_pki, 'Agent',
    )
    print(
        'mounted agents intermediate PKI at path \'{}\''
        .format(client_agents_pki)
    )

    create_policy_name = '{}-{}-agent-creation'.format(
        args.root_path,
        args.name,
    )
    client.sys.create_or_update_policy(
        name=create_policy_name,
        policy=policies.generate_agent_install_policy(args),
    )
    print('created policy {}'.format(create_policy_name))

    access_policy_name = '{}-{}-agent-access'.format(args.root_path, args.name)
    client.sys.create_or_update_policy(
        name=access_policy_name,
        policy=policies.generate_agent_access_policy(args),
    )
    print('created policy {}'.format(access_policy_name))

    res = requests.post(
        '{}/v1/auth/banana/cert/certs/{}'.format(
            os.getenv('VAULT_ADDR'),
            args.name,
        ),
        json={
            'display_name': args.name,
            'policies': '{}-{}-agent-access'.format(args.root_path, args.name),
            'certificate': agents_cert,
        },
        headers={
            'X-Vault-Token': os.getenv('VAULT_TOKEN'),
        },
        verify=not args.tls_skip_verify,
    )
    if res.status_code >= 400:
        print(res.json())
        exit(1)
    print('allowed agent certs to login into vault')

    monitor.reconfigure(args)
    print('successfully created client \'{}\''.format(args.name))


def list_clients(args):
    client = vault.get_vault_client(args)
    secrets_engines = client.sys.list_mounted_secrets_engines()
    output = []

    for key in secrets_engines['data']:
        parts = key.split('/')
        if parts[0] != 'banana' or parts[2] != 'root-pki':
            continue
        output.append([parts[1]])

    print(tabulate(output, headers=['Name']))
