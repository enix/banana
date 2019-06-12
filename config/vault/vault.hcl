ui = true
log_level = "Debug"

listener "tcp" {
  address = "0.0.0.0:8200"
	tls_cert_file = "/tls/vault.banana.enix.io.pem"
	tls_key_file = "/tls/vault.banana.enix.io.key"
}

storage "file" {
  path = "/data"
}
