import argparse
from bananadm import client
from bananadm import user
from bananadm import agent


def create(args):
    if args.type == 'client':
        client.create_client(args)
    elif args.type in ['user', 'agent']:
        if not args.client:
            print('please specify a client using --client')
            exit(1)
        if args.type == 'user':
            user.create_user(args)
        else:
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
    parser_create.add_argument(
        'name',
        nargs='?',
        help='name of the client|user to create',
    )
    parser_create.add_argument(
        '--client',
        help='client in which create the user|agent',
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
