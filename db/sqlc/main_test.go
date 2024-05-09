package db

import (
	"context"
	db_migrations "github.com/UnplugCharger/htmx-todo/migrate"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
	"testing"
	"time"
)

var testStore Store

func TestMain(m *testing.M) {
	//c, err := config.LoadConfig("../..")
	//if err != nil {
	//	log.Fatal("cannot load c:", err)
	//}
	t := &testing.T{}
	ctx := context.Background()

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatal("cannot start pg container:", err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	connPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("cannot create connection pool:", err)
	}

	migrationsPath := "../migrations"

	err = db_migrations.RunDBMigration(migrationsPath, connStr)

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
