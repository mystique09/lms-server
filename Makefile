DB_NAME=class-manager

create:
	migrate create -ext sql -dir ./database/migrations/ -seq $(name)

migrateup:
	migrate -path ./database/migrations/ -database "postgresql://mystique09:mystique09@localhost/${DB_NAME}?sslmode=disable" -verbose up

drop:
	migrate -path ./database/migrations/ -database "postgresql://mystique09:mystique09@localhost/${DB_NAME}?sslmode=disable" -verbose down

force:
	migrate -path ./database/migrations/ -database "postgresql://mystique09:mystique09@localhost/${DB_NAME}?sslmode=disable" -verbose force 1

test:
	go test -v -cover ./...

.PHONY: create migrateup drop force test