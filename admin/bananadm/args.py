import argparse
from bananadm import client
from bananadm import user
from bananadm import agent
from bananadm import input


def create(args):
    if args.type == 'client':
        args.name = input.prompt('client name')
        client.create_client(args)
    elif args.type == 'user':
        args.client = input.prompt('client in which create the user')
        args.name = input.prompt('username')
        user.create_user(args)
    elif args.type == 'agent':
        args.client = input.prompt('client in which create the agent(s)')
        agent.create_agent(args)
    else:
        print('type must be one of agent|user|client')
        exit(1)


def init_arguments():
    parser = argparse.ArgumentParser()
    subparsers = parser.add_subparsers(title='subcommands')

    parser_create = subparsers.add_parser(
        'new',
        help='create a client|user|agent'
    )
    parser_create.add_argument(
        'type',
        help='client|agent|user',
    )
    parser_create.set_defaults(func=create)

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
