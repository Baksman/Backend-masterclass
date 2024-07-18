package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/baksman/backend_masterclass/api"
	db "github.com/baksman/backend_masterclass/db/sqlc"
	"github.com/baksman/backend_masterclass/util"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://ibrahim:ibrahim@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load .env")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatalf("cannot connect to db %v", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatalf("cannot starting server %v", err)
	}

}
