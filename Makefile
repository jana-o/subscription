
postgres:
	docker run --name postgres12 -p 5433:5433 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=postgres --owner=postgres gymondo-db

dropdb:
	docker exec -it postgres12 dropdb gymondo-db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

.PHONY: postgres createdb dropdb


docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine
