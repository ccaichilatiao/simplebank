package main

import (
	"database/sql"
	"log"

	"github.com/ccaichilatiao/simplebank/api"
	db "github.com/ccaichilatiao/simplebank/db/sqlc"
	"github.com/ccaichilatiao/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
