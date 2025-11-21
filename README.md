all-inkl module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with all-inkl.

## Caddy module name

```
dns.providers.allinkl
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "allinkl",
				"kas_username": "YOUR_KAS_USERNAME",
				"kas_password": "YOUR_KAS_PASSWORD"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns allinkl {
		kas_username {env.KAS_USERNAME}
		kas_password {env.KAS_PASSWORD}
	}
}
```

```
# one site
tls {
  dns allinkl {
    kas_username {env.KAS_USERNAME}
    kas_password {env.KAS_PASSWORD}
  }
}
```
