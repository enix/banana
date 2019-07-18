# create and update policies with name matching 'banana-*'
path "sys/policy/banana-*" {
  capabilities = ["create", "update"]
}

# issue tokens for ALL policies (dangerous)
path "auth/token/create" {
  capabilities = ["sudo", "create", "update"]
}

# enable cert auth method on /auth/banana/cert
path "sys/auth/banana/cert" {
  capabilities = ["sudo", "create", "update"]
}

# set trusted CAs for login using cert method
path "auth/banana/cert/certs/*" {
  capabilities = ["create", "update"]
}

# list secret engines
path "sys/mounts" {
  capabilities = ["read"]
}

# mount new secret engines under /banana
path "sys/mounts/banana/*" {
  capabilities = ["create", "update"]
}

# manage all banana secret engines
path "banana/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}
