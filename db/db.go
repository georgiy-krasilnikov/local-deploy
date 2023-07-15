package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const defaultMigrationsDir = "../migrations"
const migrationsDirEnvName = "MIGRATIONS_DIR"

type DB struct {
	*pgxpool.Pool
}

func New(connString string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &DB{pool}, nil
}

func Migrate(connString string) error {
	var migrationsDir string
	if md := os.Getenv(migrationsDirEnvName); len(md) == 0 {
		migrationsDir = defaultMigrationsDir
	} else {
		migrationsDir = md
	}

	conn, err := goose.OpenDBWithDriver("postgres", connString)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		return err
	}

	if err := goose.Up(conn, migrationsDir); err != nil {
		return err
	}

	return nil
}
