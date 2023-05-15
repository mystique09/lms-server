DB_NAME=class-manager
DB_URL=postgresql://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable
MIGRATIONS=./database/migrations

migrate:
	migrate -path $(MIGRATIONS) -database $(DB_URL) -verbose $(op)
		
gen:
	sqlc generate

mock-store: gen
	mockgen -package mockdb -destination database/mocks/mockdb/store.go server/database/store Store
	
clean:
	rm -rf ./tmp coverage.out

lint:
	gosec -quiet -exclude-generated ./...
	gocritic check -enableAll ./...
	golangci-lint run ./...

test: clean
	go test -v -cover -coverprofile=coverage.out ./...
	
cover:
	go tool cover -html=coverage.out
	
docs:
	swagger generate spec -o ./lms-docs/swagger/swagger.json --scan-models

run:
	PORT=8000 go run cmd/main.go
	
build:
	go build -o bin/server cmd/main.go

.PHONY: migrate test run mock-store docs cover build