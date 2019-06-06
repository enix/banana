#! /usr/bin/env python3

import urllib3
from bananadm import args
from bananadm import vault


def main(argv):
    if argv.skip_tls_verify:
        urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
    if not vault.get_vault_client(argv).is_authenticated():
        print('invalid authentication data')
        exit(1)
    argv.func(argv)


if __name__ == "__main__":
    main(args.init_arguments())
