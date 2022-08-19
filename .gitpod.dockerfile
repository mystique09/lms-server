FROM gitpod/workspace-postgresql

USER gitpod
ENV PATH="$HOME/go/bin:$PATH"
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

# ENV DATABASE_URL=postgres://mystique09:mystique09@localhost/class-manager
ENV PORT=8000
ENV CLD_URL=https://
ENV FRONTEND_URL=http://localhost:3000
ENV JWT_SECRET_KEY=secretkeywthisthatlmao
ENV JWT_REFRESH_SECRET_KEY=secretrefreshkeysamplelmao
USER root