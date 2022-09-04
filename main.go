package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/zhansul19/myBank/api"
	"github.com/zhansul19/myBank/config"
	db "github.com/zhansul19/myBank/db/sqlc"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Printf("couldn't run server: %s", err.Error())
		return
	}
	if err := server.Run(config.ServerAddress); err != nil {
		log.Fatal(err)
	}
}
