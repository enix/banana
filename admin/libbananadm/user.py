import os
import sys
from libbananadm import vault
from tabulate import tabulate


def create_user(args):
    client_pki = '{}/{}/users-pki'.format(args.root_path, args.client)
    sys.stderr.write(
        'issuing certificate with CN \'{}\' using PKI {}\n'
        .format(args.name, client_pki)
    )
    cert, key = vault.create_cert(
        args, args.name, client_pki, 'default'
    )
    open(args.name + '.pem', 'w').write(cert)
    open(args.name + '.key', 'w').write(key)

    sys.stderr.write('creating p12 file\n')
    os.system(
        'openssl pkcs12 -export -inkey {}.key -in {}.pem'
        .format(args.name, args.name, args.name)
    )
    os.system('rm {}.pem'.format(args.name))
    os.system('rm {}.key'.format(args.name))

    sys.stderr.write('successfully wrote p12 data to stdout\n')


def list_users(args):
    mount_point = '{}/{}/users-pki'.format(args.root_path, args.client)
    users = filter(
        lambda x: x[0] != '{} User Intermediate CA'.format(args.client),
        vault.list_cn_from_pki(args, mount_point),
    )

    print(tabulate(users, headers=['Name', 'Serial']))
