DB_NAME=class-manager

setup:
	go install github.com/cosmtrek/air@latest \
	# install sqlc
	 \
	# install golang-migrate
	go get github.com/golang-migrate/migrate@latest \
	# install package dependencies
	go install

dev:
	cd frontend && pnpm watch & go run cmd/main.go && fg

run: 
    docker run --name mystique09 -e POSTGRES_PASSWORD=mystique09 -d postgres

watch:
	cd frontend && pnpm build & air && fg

create:
	migrate create -ext sql -dir ./database/migrations/ -seq $(name)

migrateup:
	migrate -path ./database/migrations/ -database "postgresql://mystique09:mystique09@localhost/${DB_NAME}?sslmode=disable" -verbose up

drop:
	migrate -path ./database/migrations/ -database "postgresql://mystique09:mystique09@localhost/${DB_NAME}?sslmode=disable" -verbose down

force:
	migrate -path ./database/migrations/ -database "postgresql://mystique09:mystique09@localhost/${DB_NAME}?sslmode=disable" -verbose force 1

.PHONY: create migrateup drop force setup run
