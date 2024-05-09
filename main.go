package main

import (
	"context"
	"github.com/UnplugCharger/htmx-todo/api"
	"github.com/UnplugCharger/htmx-todo/config"
	db "github.com/UnplugCharger/htmx-todo/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func main() {
	// Start the server
	ctx := context.Background()
	// Load the configuration
	cnf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	conPool, err := pgxpool.New(ctx, cnf.DBSource)

	store := db.NewStore(conPool)

	server, err := api.NewServer(store, cnf)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create server")
	}

	err = server.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

}
