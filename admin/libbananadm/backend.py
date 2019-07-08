import os
from libbananadm import input
from libbananadm import vault
from tabulate import tabulate


def prompt_s3_secret_values():
    return {
        'AWS_ACCESS_KEY_ID': input.prompt(
            'AWS_ACCESS_KEY_ID', os.getenv('AWS_ACCESS_KEY_ID'),
        ),
        'AWS_SECRET_ACCESS_KEY': input.prompt(
            'AWS_SECRET_ACCESS_KEY', os.getenv('AWS_SECRET_ACCESS_KEY'),
        ),
    }


def prompt_secret_values(type):
    return {
        's3': prompt_s3_secret_values,
    }[type]()


def create_backend_secret(args, values):
    client = vault.get_vault_client(args)
    secret_path = 'backends/{}'.format(args.name)
    mount_point = '{}/{}/secrets'.format(args.root_path, args.client)
    client.secrets.kv.v2.create_or_update_secret(
        path=secret_path,
        mount_point=mount_point,
        secret=values,
    )

    print(
        'successfully saved secret {} in KV engine {}'
        .format(secret_path, mount_point)
    )


def list_backends(args):
    client = vault.get_vault_client(args)
    mount_point = '{}/{}/secrets'.format(args.root_path, args.client)
    secret_list = client.secrets.kv.v2.list_secrets(
        'backends',
        mount_point=mount_point,
    )
    output = map(lambda k: [k], secret_list['data']['keys'])
    # output = []

    # for secret in secret_list['data']['keys']:
    #     secret_data = client.secrets.kv.v2.read_secret_version(
    #         path='backends/{}'.format(secret),
    #         mount_point=mount_point,
    #     )

    #     output.append([secret, 's3'])

    print(tabulate(output, headers=['Name']))
