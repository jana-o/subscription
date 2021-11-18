
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root gymondo-db

dropdb:
	docker exec -it postgres12 dropdb gymondo-db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gymondo-db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gymondo-db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

test:
	@ echo "-> start tests..."
	@ go test -v -cover ./...
	@ echo "-> done!"

lint-go:
	@ echo "-> Running linters ..."
	@ golangci-lint run --print-resources-usage --exclude-use-default=false
	@ echo "-> Done!"

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server test lint-go

docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine
