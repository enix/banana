import argparse
from bananadm import client
from bananadm import user
from bananadm import agent
from bananadm import backend
from bananadm import monitor
from bananadm import input


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
        help='setup monitor policies and get its certificate'
    ).set_defaults(func=monitor.init)

    subparsers.add_parser(
        'reconfigure',
        help='ask monitor\'s nginx to re-fetch all CA certs'
    ).set_defaults(func=reconfigure)

    parser_create = subparsers.add_parser(
        'new',
        help='create a client|user|agent'
    )
    subparsers_create = parser_create.add_subparsers(
        title='types that can be created',
    )

    subparsers_create.add_parser(
        'client',
        help='create a client',
    ).set_defaults(func=create_client)
    subparsers_create.add_parser(
        'user',
        help='create a user',
    ).set_defaults(func=create_user)
    subparsers_create.add_parser(
        'agent',
        help='create a agent',
    ).set_defaults(func=create_agent)

    parser_backend_secret = subparsers_create.add_parser(
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
        '--skip-tls-verify',
        help='ignore TLS errors',
        action='store_true'
    )
    parser.add_argument(
        '--root-path',
        help='root vault mount point',
        default='banana',
    )

    return parser.parse_args()
