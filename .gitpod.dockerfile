FROM gitpod/workspace-postgresql

USER gitpod
ENV PATH="$HOME/go/bin:$PATH"
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

# Installing golang-migrate
RUN echo "Installing golang-migrate/migrate ..."
RUN go get -u -d github.com/golang-migrate/migrate/cmd/migrate
RUN cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
RUN git checkout master  # e.g. v4.1.0
# Go 1.16+
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$TAG

# Create new database user
# ENV DATABASE_URL=postgres://mystique09:@localhost/class-manager
ENV PORT=8000
ENV CLD_URL=https://
ENV FRONTEND_URL=http://localhost:3000
ENV JWT_SECRET_KEY=secretkeywthisthatlmao
ENV JWT_REFRESH_SECRET_KEY=secretrefreshkeysamplelmao

USER root
