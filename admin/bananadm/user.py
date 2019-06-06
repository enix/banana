import os
from bananadm import vault


def create_user(args):
    client_pki = '{}/{}/users-pki'.format(args.root_path, args.client)
    cert, key = vault.create_cert(
        args, args.name, client_pki, 'default'
    )
    open(args.name + '.pem', 'w').write(cert)
    open(args.name + '.key', 'w').write(key)
    os.system(
        'openssl pkcs12 -export -out {}.p12 -inkey {}.key -in {}.pem'
        .format(args.name, args.name, args.name)
    )
    os.system('rm {}.pem'.format(args.name))
    os.system('rm {}.key'.format(args.name))
