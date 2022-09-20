package main

import (
	"database/sql"
	"github.com/RakibSiddiquee/simplebank/api"
	db "github.com/RakibSiddiquee/simplebank/db/sqlc"
	"github.com/RakibSiddiquee/simplebank/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".") // dot is used as app.env file and main.go file in the same directory
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
