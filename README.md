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

## Docker example

1) Install docker and docker-compose, then run:
```bash
apt install docker.io docker-compose
```

2) Build the docker image using the provided Dockerfile and build script:

```bash
cd test-docker
chmod +x build/build.sh
./build/build.sh
```

3) Change the name of the `.env.example` to `.env` file to include your all-inkl KAS credentials:


4) Edit the `Caddyfile` in the `test-docker/caddy` directory to configure your desired domains and settings.

5) Start the Caddy server with docker-compose:

```bash
docker-compose up -d
```