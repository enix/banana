agent_install_policy_template = '''
path "{root_path}/{company}/agents-pki/issue/default" {{
    capabilities = ["create", "update"]
}}
'''

agent_access_policy_template = '''
path "{root_path}/{company}/secrets/data/backends/*" {{
    capabilities = ["read"]
}}

path "{root_path}/{company}/secrets/data/agents/{{{{identity.entity.id}}}}"\
{{
    capabilities = ["read", "list", "create", "update"]
}}
'''

monitor_policy_template = '''
path "sys/mounts" {{
  capabilities = ["read"]
}}

path "{root_path}/+/root-pki/ca/pem" {{
  capabilities = ["read"]
}}

path "{root_path}/+/agents-pki/ca/pem" {{
  capabilities = ["read"]
}}

path "{root_path}/+/users-pki/ca/pem" {{
  capabilities = ["read"]
}}
'''


def generate_agent_install_policy(args):
    return agent_install_policy_template.format(
        root_path=args.root_path,
        company=args.name,
    )


def generate_agent_access_policy(args):
    return agent_access_policy_template.format(
        root_path=args.root_path,
        company=args.name,
    )


def generate_monitor_policy(args):
    return monitor_policy_template.format(
        root_path=args.root_path,
    )
