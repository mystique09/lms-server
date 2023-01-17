# LMS Server
Back-end for my [LMS Website](https://class-management.vercel.app)  website. 

> **Warning**
> Website is under development, I need to finish this server first before I refactor the client/frontend.

#### Built with
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)

## Getting started
LMS server require this binaries to be present in order to run.

### Prerequisites
Install required tools if you want to self-host.
- [Go](https://go.dev/dl) version ^1.19
- PostgreSQL

### Installation
1. Clone repository
	```bash
	git clone https://github.com/mystique09/lms-server
	```
1. Install dependencies
	```bash
	cd lms-server
	go mod tidy
	```
1. Copy the app.sample.env to app.env add your environment variables
	```bash
	cp app.sample.env app.env
	# or
	cat app.sample.env > app.env
	```
1. Run the server
	```bash
	go run cmd/main.go
	# or, if you have cmake/make
	make run
	```

### Development
While developing, you need to install this binaries.
1. Sqlc, sqlc is used to generate typesafe sql queries. [Install sqlc](https://docs.sqlc.dev/en/latest/overview/install.html).
1. golang-migrate for manual migration. [Install golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).
1. mockgen for code mocks, also used for testing. [Install gomock](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).
1. Gosec, a go linter for security. [Install gosec](https://github.com/securego/gosec#go-116).
1. Go-critic, a linter used to critic go code. [Install go-critic](https://github.com/go-critic/go-critic#installation).
1. Golangci-lint, powerful golang linter. [Install golangci-lint](https://golangci-lint.run/usage/install/).

To run tests:
```bash
	go test -v -cover -coverprofile=coverage.out
	# or
	make test
```
To run linters:
```bash
	gosec -quiet -exclude-generated ./...
	gocritic check -enableAll ./...
	golangci-lint run ./...
	# or
	make lint
```

[Apache License 2.0](./LICENSE)