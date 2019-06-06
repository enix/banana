import argparse
from bananadm import client
from bananadm import user
from bananadm import agent
from bananadm import input


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


def init_arguments():
    parser = argparse.ArgumentParser()
    subparsers = parser.add_subparsers(title='subcommands')

    parser_create = subparsers.add_parser(
        'new',
        help='create a client|user|agent'
    )
    subparsers_create = parser_create.add_subparsers(
        title='types that can be created',
    )
    parser_client = subparsers_create.add_parser(
        'client',
        help='create a client',
    )
    parser_client.set_defaults(func=create_client)
    parser_user = subparsers_create.add_parser(
        'user',
        help='create a user',
    )
    parser_user.set_defaults(func=create_user)
    parser_agent = subparsers_create.add_parser(
        'agent',
        help='create a agent',
    )
    parser_agent.set_defaults(func=create_agent)

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
