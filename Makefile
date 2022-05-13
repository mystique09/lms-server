create:
	migrate create -ext sql -dir ./db/migrations/ -seq $(name)

migrateup:
	migrate -path ./db/migrations/ -database "postgresql://mystique09:mystique09@localhost/class-management?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migrations/ -database "postgresql://mystique09:mystique09@localhost/class-management?sslmode=disable" -verbose down

force:
	migrate -path ./db/migrations/ -database "postgresql://mystique09:mystique09@localhost/class-management?sslmode=disable" -verbose force 1

.PHONY: create migrate reset force
