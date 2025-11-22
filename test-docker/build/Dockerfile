# Use Caddy's official builder image as the base
FROM caddy:2.10.2-builder-alpine as builder

# Install additional dependencies
RUN apk add --no-cache \
    bash \
    curl \
    gettext \
    ca-certificates \
    tini

RUN xcaddy build && rm ./caddy 

# Use xcaddy to build a custom Caddy with the necessary plugins
RUN xcaddy build \
    --with github.com/libdns/all-inkl \
    --with github.com/caddy-dns/all-inkl

# Use the official Caddy image as the base for the final image
FROM caddy:2.10.2

# Copy the built Caddy binary into the final image
COPY --from=builder /usr/bin/caddy /usr/bin/caddy