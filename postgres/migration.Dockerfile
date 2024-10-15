FROM alpine:3.20

# Install bash and clean up
RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

# Download Goose binary for migrations
ADD https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

# Set working directory
WORKDIR /root

# Add migration files and scripts
ADD migrations/*.sql migrations/
ADD migration.sh .
ADD .env .

# Ensure migration_prod.sh is added if required
# ADD migration_prod.sh .  # Uncomment or add this if the script is needed

# Make migration.sh executable
RUN chmod +x migration.sh

# Set entrypoint to execute migration.sh
ENTRYPOINT ["bash", "migration.sh"]
