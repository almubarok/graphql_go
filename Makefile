migrate:
	@go run cmd/migration/migration.go

build:
	@go build -o bin

run: build
	@./bin

mocks:
	@mockery --all --keeptree --output=./mocks --case underscore