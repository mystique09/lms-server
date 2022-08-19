FROM gitpod/workspace-full:latest

USER root
RUN apt-get update && apt-get install -y \
        postgresql \
        postgresql-contrib \
    && apt-get clean && rm -rf /var/cache/apt/* && rm -rf /var/lib/apt/lists/* && rm -rf /tmp/*

# Setup postgres for gitpod workspace
USER gitpod
ENV PATH="/usr/lib/postgresql/10/bin:$PATH"
RUN mkdir -p ~/pg/data; mkdir -p ~/pg/scripts; mkdir -p ~/pg/logs; mkdir -p ~/pg/sockets; initdb -D ~/pg/data/

# Scripts
# start
RUN '#!/bin/bash\n\
pg_ctl -D ~/pg/data/ -l ~/pg/logs/log -o "-k ~/pg/sockets" start' > ~/pg/scripts/pg_start.sh
# stop
RUN '#!/bin/bash\n\
pg_ctl -D ~/pg/data/ -l ~/pg/logs/log -o "-k ~/pg/sockets" stop' > ~/pg/scripts/pg_stop.sh

RUN chmod +x ~/pg/scripts/*
ENV PATH="$HOME/pg/scripts:$PATH"

ENV PATH="$HOME/go/bin:$PATH"
RUN go install github.com/golang-migrate/migrate@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

ENV DATABASE_URL=postgres://mystique09:mystique09@localhost/class-manager
ENV PORT=8000
ENV CLD_URL=https://

USER root
