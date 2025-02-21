package main

import (
	"context"
	"log"

	"github.com/Abenezer0923/simple-bank/api"
	db "github.com/Abenezer0923/simple-bank/db/sqlc"
	"github.com/Abenezer0923/simple-bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
	
)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load the config:", err)
	}
	// Create a connection pool using pgxpool
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer connPool.Close()

	// Create a new store with the connection pool
	store := db.NewStore(connPool)

	// Create a new server with the store
	server := api.NewServer(store)

	// Start the server
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}