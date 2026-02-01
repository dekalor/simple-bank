package main

import (
	"database/sql"
	"log"

	"github.com/dekalor/simple-bank/api"
	db "github.com/dekalor/simple-bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver       = "postgres"
	dbSource       = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
	serverrAddress = "localhost:3000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverrAddress)
	if err != nil {
		log.Fatal("Cannot start the server", err)
	}
}
