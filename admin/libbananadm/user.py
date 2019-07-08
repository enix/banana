import os
from libbananadm import vault
from cryptography import x509, hazmat


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
    client = vault.get_vault_client(args)
    mount_point = '{}/{}/users-pki'.format(args.root_path, args.client)
    response = client.secrets.pki.list_certificates(
        mount_point=mount_point,
    )
    for key in response['data']['keys']:
        user_cert = client.secrets.pki.read_certificate(
            serial=key,
            mount_point=mount_point,
        )

        cert_str = user_cert['data']['certificate']
        cert = x509.load_pem_x509_certificate(
            cert_str.encode('ascii'),
            hazmat.backends.default_backend(),
        )

        name = cert.subject.get_attributes_for_oid(
            x509.NameOID.COMMON_NAME,
        )[0].value

        if name != args.client + ' User Intermediate CA':
            print('*', name)
