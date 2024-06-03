postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

removepostgres:
	docker rm -f postgres

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up ${version}

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down ${version}

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go build -o main main.go && ./main

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ccaichilatiao/simplebank/db/sqlc Store

.PHONY: postgres removepostgres createdb dropdb migrateup migratedown sqlc test server