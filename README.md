# subscription

network:
docker network create my-network

postgres:
docker run --name postgres12 --network my-network -p 5433:5433 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres:12-alpine
// change back to 5432!!!

migrateup:
migrate -path db/migration -database "postgresql://postgres:password@localhost:5433/gymondo?sslmode=disable" -verbose up
 // change back to 5432!!!

run app: 
    go run main.go

test:
    go test

mock: 

migrate


Create my-network, start postgres container, create db, run db with migrateup run app and test