DB_NAME=class-manager

clean:
	rm -rf ./tmp coverage.out

setup:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1 \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz \
    sudo mv migrate /usr/bin

lint:
	golangci-lint run ./...

test: clean
	go test -v -cover -coverprofile=coverage.out ./...

server:
	go run cmd/main.go

postgres:
	sudo docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root class-manager

dropdb:
	docker exec -it postgres dropdb class-manager

migrateup:
	migrate -path ./database/migrations/ -database "postgresql://root:secret@localhost:5432/class-manager?sslmode=disable" -verbose up

migratedown:
	migrate -path ./database/migrations/ -database "postgresql://root:secret@localhost:5432/class-manager?sslmode=disable" -verbose down

force:
	migrate -path ./database/migrations/ -database "postgresql://root:secret@localhost:5432/class-manager?sslmode=disable" -verbose force 1

sqlc:
	sqlc generate

mock: sqlc
	mockgen -package mockdb -destination database/mock/store.go server/database/sqlc Store

.PHONY: postgres create migrateup migratedown force test server mock