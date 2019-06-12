import hvac
import os


def get_vault_client(args):
    verify = False if args.skip_tls_verify else True
    vault = hvac.Client(
        url=os.getenv('VAULT_ADDR'),
        token=os.getenv('VAULT_TOKEN'),
        verify=verify,
    )
    return vault


def create_cert(args, cn, pki, role):
    vault = get_vault_client(args)
    res = vault.secrets.pki.generate_certificate(
        name=role,
        common_name=cn,
        mount_point=pki,
    )
    cert = res.json()['data']['certificate']
    key = res.json()['data']['private_key']
    return cert, key


def create_intermediate_ca(args, root_path, int_path, type):
    vault = get_vault_client(args)
    vault.sys.enable_secrets_engine('pki', path=int_path, config={
        'max_lease_ttl': '43800h',
    })
    users_int = vault.secrets.pki.generate_intermediate(
        type='internal',
        common_name='{} {} Intermediate CA'.format(args.name, type),
        mount_point=int_path,
        extra_params={
            'ttl': '43800h',
        },
    )
    signed_users_int = vault.secrets.pki.sign_intermediate(
        csr=users_int.json()['data']['csr'],
        common_name='{} {} Intermediate CA'.format(args.name, type),
        mount_point=root_path,
        extra_params={
            'ttl': '43800h',
        },
    )
    int_cert = signed_users_int.json()['data']['certificate']
    vault.secrets.pki.set_signed_intermediate(
        int_cert,
        mount_point=int_path,
    )
    vault.secrets.pki.create_or_update_role('default', {
        'allow_any_name': 'true',
        'organization': args.name,
        'ou': type,
        'default_lease_ttl': '17520h',
    }, mount_point=int_path)
    return int_cert
