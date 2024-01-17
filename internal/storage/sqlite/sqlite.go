package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func New(storagePath string, migrations []string) (*sql.DB, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = executeMigrations(db, migrations)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}

func executeMigrations(db *sql.DB, migrations []string) error {
	transaction, err := db.Begin()
	if err != nil {
		return fmt.Errorf("unable to start transaction: %w", err)
	}

	for _, migration := range migrations {
		stmt, err := transaction.Prepare(migration)
		if err != nil {
			transaction.Rollback()
			return fmt.Errorf("unable to prepare migration: %w", err)
		}

		_, err = stmt.Exec()
		if err != nil {
			transaction.Rollback()
			return fmt.Errorf("unable to execute migration: %w", err)
		}
	}

	err = transaction.Commit()
	if err != nil {
		return fmt.Errorf("unable to execute all migrations: %w", err)
	}

	return nil
}
