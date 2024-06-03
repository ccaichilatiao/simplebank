package main

import (
	"database/sql"
	"log"

	"github.com/ccaichilatiao/simplebank/api"
	db "github.com/ccaichilatiao/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

var (
	dbDriver      = "postgres"
	dbResource    = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbResource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
