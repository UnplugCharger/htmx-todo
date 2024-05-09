package db_migrations

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"
	"os"
)

func RunDBMigration(migrationURL string, dbSource string) error {
	log.Info().Msgf("migration url: %s", migrationURL)

	fPath, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot get current working directory")
	}
	log.Info().Msgf("current working directory: %s", fPath)
	migrationURL = "file://" + migrationURL
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && !errors.Is(migrate.ErrNoChange, err) {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
	return nil
}
