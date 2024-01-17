package storage

import (
	"database/sql"
	"fmt"

	"github.com/anesthesia-band/golang-shop-api/internal/storage/sqlite"
	"github.com/anesthesia-band/golang-shop-api/storage/migrations"
)

type Storage struct {
	DB *sql.DB
}

func Init(storagePath string) (*Storage, error) {
	db, err := sqlite.New(storagePath, migrations.Migrations)
	if err != nil {
		return nil, fmt.Errorf("unable to init storage: %w", err)
	}

	return &Storage{DB: db}, nil
}
