import os
from libbananadm import vault
from tabulate import tabulate


def create_user(args):
    client_pki = '{}/{}/users-pki'.format(args.root_path, args.client)
    print(
        'issuing certificate with CN \'{}\' using PKI {}'
        .format(args.name, client_pki)
    )
    cert, key = vault.create_cert(
        args, args.name, client_pki, 'default'
    )
    open(args.name + '.pem', 'w').write(cert)
    open(args.name + '.key', 'w').write(key)

    print('creating p12 file')
    os.system(
        'openssl pkcs12 -export -out {}.p12 -inkey {}.key -in {}.pem'
        .format(args.name, args.name, args.name)
    )
    os.system('rm {}.pem'.format(args.name))
    os.system('rm {}.key'.format(args.name))

    print('successfully wrote {}.p12'.format(args.name))


def list_users(args):
    mount_point = '{}/{}/users-pki'.format(args.root_path, args.client)
    users = filter(
        lambda x: x[0] != '{} User Intermediate CA'.format(args.client),
        vault.list_cn_from_pki(args, mount_point),
    )

    print(tabulate(users, headers=['Name', 'Serial']))
