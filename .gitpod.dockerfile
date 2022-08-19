FROM gitpod/workspace-postgresql

USER gitpod
ENV PATH="$HOME/go/bin:$PATH"
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

RUN apt-get update && \ 
     apt-get install -y \ 
     apt-transport-https \ 
     ca-certificates \ 
     curl \
     gnupg-agent

# Installing golang-migrate
RUN curl -sSL https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
RUN echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
RUN apt-get update
RUN apt-get install -y migrate

# Create new database user

# ENV DATABASE_URL=postgres://mystique09:@localhost/class-manager
ENV PORT=8000
ENV CLD_URL=https://
ENV FRONTEND_URL=http://localhost:3000
ENV JWT_SECRET_KEY=secretkeywthisthatlmao
ENV JWT_REFRESH_SECRET_KEY=secretrefreshkeysamplelmao

USER root
