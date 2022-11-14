DB_NAME=class-manager

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root class-manager

dropdb:
	docker exec -it postgres12 dropdb class-manager

migrateup:
	migrate -path ./database/migrations/ -database "postgresql://root:secret@localhost/${DB_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path ./database/migrations/ -database "postgresql://root:secret@localhost/${DB_NAME}?sslmode=disable" -verbose down

force:
	migrate -path ./database/migrations/ -database "postgresql://root:secret@localhost/${DB_NAME}?sslmode=disable" -verbose force 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres create migrateup migratedown force test