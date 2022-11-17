DB_NAME=class-manager

clean:
	rm -rf ./tmp coverage.out

lint:
	golangci-lint run ./...

security:
	gosec -quiet -exclude-dir database/sqlc ./...

critic:
	gocritic check -enableAll ./...

test: clean lint security critic
	go test -v -cover ./...

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