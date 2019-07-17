import argparse
from libbananadm import client
from libbananadm import user
from libbananadm import agent
from libbananadm import backend
from libbananadm import monitor
from libbananadm import input


def reconfigure(args):
    monitor.reconfigure(args)


def create_user(args):
    args.client = input.prompt('client in which create the user')
    args.name = input.prompt('username')
    user.create_user(args)


def create_agent(args):
    args.client = input.prompt('client in which create the agent(s)')
    agent.create_agent(args)


def create_client(args):
    args.name = input.prompt('client name')
    client.create_client(args)


def create_backend_secret(args, type):
    args.client = input.prompt('client in which create the secret')
    args.name = input.prompt('backend name')
    values = backend.prompt_secret_values(type)
    backend.create_backend_secret(args, values)


def init_arguments():
    parser = argparse.ArgumentParser()
    subparsers = parser.add_subparsers(title='subcommands')

    subparsers.add_parser(
        'init',
        help='setup monitor policies'
    ).set_defaults(func=monitor.init)

    subparsers.add_parser(
        'reconfigure',
        help='ask monitor\'s nginx to re-fetch all CA certs'
    ).set_defaults(func=reconfigure)

    parser_list = subparsers.add_parser(
        'list',
        help='list clients|users|agents|backends'
    )
    subparsers_list = parser_list.add_subparsers(
        title='types of resources',
    )

    subparsers_list.add_parser(
        'clients',
        help='list clients stored in vault',
    ).set_defaults(func=client.list_clients)

    list_users = subparsers_list.add_parser(
        'users',
        help='list users for a given client',
    )
    list_users.add_argument('client', help='client name')
    list_users.set_defaults(func=user.list_users)

    list_agents = subparsers_list.add_parser(
        'agents',
        help='list agents for a given client',
    )
    list_agents.add_argument('client', help='client name')
    list_agents.set_defaults(func=agent.list_agents)

    list_backends = subparsers_list.add_parser(
        'backends',
        help='list backends for a given client',
    )
    list_backends.add_argument('client', help='client name')
    list_backends.set_defaults(func=backend.list_backends)

    parser_new = subparsers.add_parser(
        'new',
        help='create a client|user|agent|backend'
    )
    subparsers_new = parser_new.add_subparsers(
        title='types that can be created',
    )

    subparsers_new.add_parser(
        'client',
        help='create a client',
    ).set_defaults(func=create_client)
    subparsers_new.add_parser(
        'user',
        help='create a user',
    ).set_defaults(func=create_user)
    subparsers_new.add_parser(
        'agent',
        help='create a agent',
    ).set_defaults(func=create_agent)

    parser_backend_secret = subparsers_new.add_parser(
        'backend',
        help='add a backend secret',
    )
    subparsers_backend_secret = parser_backend_secret.add_subparsers(
        title='supported backend types',
    )
    subparsers_backend_secret.add_parser(
        's3',
        help='s3 protocol credentials, also works \
            with openstack block storage',
    ).set_defaults(func=lambda args: create_backend_secret(args, 's3'))

    parser.add_argument(
        '--tls-skip-verify',
        help='ignore TLS errors',
        action='store_true'
    )
    parser.add_argument(
        '--root-path',
        help='root vault mount point',
        default='banana',
    )
    parser.add_argument(
        '--monitor-addr',
        help='monitor API address',
        default='https://api.banana.enix.io',
    )

    return parser
