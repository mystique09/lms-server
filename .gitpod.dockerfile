FROM gitpod/workspace-postgresql

USER gitpod
ENV PATH="$HOME/go/bin:$PATH"
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

ENV DATABASE_URL=postgres://mystique09:mystique09@localhost/class-manager
ENV PORT=8000
ENV CLD_URL=https://

USER root
