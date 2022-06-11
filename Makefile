create:
	migrate create -ext sql -dir ./database/migrations/ -seq $(name)

migrateup:
	migrate -path ./database/migrations/ -database "postgresql://mystique09:mystique09@localhost/class-management?sslmode=disable" -verbose up

migratedown:
	migrate -path ./database/migrations/ -database "postgresql://mystique09:mystique09@localhost/class-management?sslmode=disable" -verbose down

force:
	migrate -path ./database/migrations/ -database "postgresql://mystique09:mystique09@localhost/class-management?sslmode=disable" -verbose force 1

.PHONY: create migrate reset force
