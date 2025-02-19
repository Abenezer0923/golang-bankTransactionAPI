package main

import (
	"context"
	"log"

	"github.com/Abenezer0923/simple-bank/api"
	db "github.com/Abenezer0923/simple-bank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource      = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	// Create a connection pool using pgxpool
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer connPool.Close()

	// Create a new store with the connection pool
	store := db.NewStore(connPool)

	// Create a new server with the store
	server := api.NewServer(store)

	// Start the server
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}
