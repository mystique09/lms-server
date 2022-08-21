#syntax=docker/dockerfile:1
FROM go:latest
RUN apk add --no-cache make
WORKDIR /app
COPY . .

ENV PORT=8080

RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN go install go get github.com/golang-migrate/migrate@latest
RUN go get && go build cmd/main.go

RUN ./main
EXPOSE ${PORT}
