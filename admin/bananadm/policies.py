agent_install_policy_template = '''
path "{root_path}/{company}/agents-pki/issue/default" {{
    capabilities = ["create", "update"]
}}

path "{root_path}/{company}/secrets/agents/{{{{identity.entity.name}}}}" {{
    capabilities = ["create", "update"]
}}
'''

agent_access_policy_template = '''
path "{root_path}/{company}/secrets/backends/*" {{
    capabilities = ["read"]
}}

path "{root_path}/{company}/secrets/agents/{{{{identity.entity.name}}}}/*" {{
    capabilities = ["read", "list", "create", "update"]
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
