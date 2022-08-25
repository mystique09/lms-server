FROM gitpod/workspace-postgresql

USER gitpod
ENV PATH="$HOME/go/bin:$PATH"
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

# Create new database user
# RUN createuser --host localhost -d -l -s mystique09
# RUN createdb --host=localhost class-manager

# Server config environment variables
ENV DATABASE_URL=postgres://gitpod@localhost/postgres?sslmode=disable
ENV PORT=8000
ENV CLD_URL=https://
ENV FRONTEND_URL=http://localhost:3000
ENV JWT_SECRET_KEY=secretkeywthisthatlmao
ENV JWT_REFRESH_SECRET_KEY=secretrefreshkeysamplelmao

USER root
