package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/zhansul19/myBank/api"
	db "github.com/zhansul19/myBank/db/sqlc"
	"github.com/zhansul19/myBank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	runDBMigration(config.MigrationURL, config.DBSource)

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
func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated successfully")
}
