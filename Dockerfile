#syntax=docker/dockerfile:1
FROM alpine:latest

RUN apk add --no-cache make git musl-dev go

ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

WORKDIR /app
COPY . .

ENV PORT=8080

RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN go install go get github.com/golang-migrate/migrate@latest

RUN go get && go build cmd/main.go

EXPOSE ${PORT}
RUN ./main