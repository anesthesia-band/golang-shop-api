package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/anesthesia-band/golang-shop-api/internal/storage"
	_ "github.com/mattn/go-sqlite3"
)

func New(storagePath string) (*storage.Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(migration)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare migration: %w", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("unable to complete migration: %w", err)
	}

	return &storage.Storage{DB: db}, nil
}

// TODO: Add migrations
const migration string = `
	create table if not exists goods (
		id integer not null primary key autoincrement,
		name varchar(255) not null,
		type varchar(255) not null,
		data TEXT not null,
		active BOOLEAN not null default true,
		created_at datetime not null default CURRENT_TIMESTAMP,
		updated_at datetime not null default CURRENT_TIMESTAMP
  	)

	create index if not exists idx__goods__type on goods(type)

	create table if not exists groups (
		id integer not null primary key autoincrement,
		name varchar(255) not null,
		type varchar(255) not null,
		data TEXT not null,
		active BOOLEAN not null default true,
		created_at datetime not null default CURRENT_TIMESTAMP,
		updated_at datetime not null default CURRENT_TIMESTAMP
	)

	create index if not exists idx__groups__type on groups(type)

	create table if not exists groups_goods (
		group_id INTEGER not null,
		good_id INTEGER not null,
		active BOOLEAN not null default true,
		created_at datetime not null default CURRENT_TIMESTAMP,
		updated_at datetime not null default CURRENT_TIMESTAMP,

		primary key (kit_id, good_id)
	)

	create table orders (
		id integer not null primary key autoincrement,
		name varchar(255) null,
		phone varchar(255) not null,
		email varchar(255) null,
		order_data TEXT null,
		status VARCHAR(255) not null default 'new',
		created_at datetime not null default CURRENT_TIMESTAMP,
		updated_at datetime not null default CURRENT_TIMESTAMP
	)

	create index if not exists idx__orders__phone on orders(phone)
	create index if not exists idx__orders__status on orders(status)
`
