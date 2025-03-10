// main.go
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

	// Use pgxpool instead of sql.Open
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer connPool.Close()

	store := db.NewStore(connPool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}