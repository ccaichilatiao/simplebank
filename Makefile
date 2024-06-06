DB_URL=postgresql://root:password@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

removepostgres:
	docker rm -f postgres

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up ${version}

migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down ${version}

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go build -o main main.go && ./main

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ccaichilatiao/simplebank/db/sqlc Store

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

proto:
	rm -f pb/*.go
	protoc --proto_path=proto \
	--go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: postgres removepostgres createdb dropdb migrateup migratedown sqlc test server db_schema proto