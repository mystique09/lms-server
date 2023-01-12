DB_NAME=class-manager

pgstart:
	sudo docker run --name pg -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

pgstop:
	sudo docker stop $(id)

pgrestart:
	sudo docker container restart pg

createdb:
	sudo docker exec -it pg createdb --username=root --owner=root $(DB_NAME)

dropdb:
	sudo docker exec -it pg dropdb $(DB_NAME)

migrateup:
	migrate -path ./database/migrations/ -database "postgresql://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path ./database/migrations/ -database "postgresql://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down

force:
	migrate -path ./database/migrations/ -database "postgresql://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose force 1

inspect:
	sudo docker exec -it pg psql $(DB_NAME)

sqlc:
	sqlc generate

mock: sqlc
	mockgen -package mockdb -destination database/mocks/mockdb/store.go server/database/store Store
	
clean:
	rm -rf ./tmp coverage.out

setup:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1 \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz \
    sudo mv migrate /usr/bin

lint:
	gosec -quiet -exclude-generated ./...
	gocritic check -enableAll ./...
	golangci-lint run ./...

test: clean
	go test -v -cover -coverprofile=coverage.out ./...
	
cover:
	go tool cover -html=coverage.out
	
docs:
	rm -rf ./lms-docs/swagger/swagger.json
	swagger generate spec -o ./lms-docs/swagger/swagger.json --scan-models

run:
	go run cmd/main.go
	
build:
	rm -rf lms-server
	go build -o lms-server cmd/main.go

temp-server: build
	rm -rfd test-server
	mkdir -p test-server/database/migrations
	cp database/migrations/* test-server/database/migrations
	cp app.env test-server
	mv lms-server test-server

.PHONY: postgres create migrateup migratedown force test run mock docs cover build temp-server