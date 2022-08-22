package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/zhansul19/myBank/api"
	"github.com/zhansul19/myBank/db"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	address  = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)
	if err := server.Run(address); err != nil {
		log.Fatal(err)
	}
}
