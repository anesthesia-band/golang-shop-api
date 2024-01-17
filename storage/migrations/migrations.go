package migrations

var Migrations = []string{
	`create table if not exists goods (
		id integer not null primary key autoincrement,
		name varchar(255) not null,
		type varchar(255) not null,
		data TEXT not null,
		price varchar(255) not null,
		active BOOLEAN not null default true
  	)`,
	`create index if not exists idx__goods__type on goods(type)`,
	`create table if not exists groups (
		id integer not null primary key autoincrement,
		name varchar(255) not null,
		type varchar(255) not null,
		data TEXT not null,
		price varchar(255) not null,
		active BOOLEAN not null default true
	)`,
	`create index if not exists idx__groups__type on groups(type)`,
	`create table if not exists groups_goods (
		group_id INTEGER not null,
		good_id INTEGER not null,
		quantity INTEGER not null,

		primary key (group_id, good_id)
	)`,
	`create table if not exists orders (
		id integer not null primary key autoincrement,
		name varchar(255) null,
		phone varchar(255) not null,
		email varchar(255) null,
		order_data TEXT null,
		status VARCHAR(255) not null default 'new'
	)`,
	`create index if not exists idx__orders__phone on orders(phone)`,
	`create index if not exists idx__orders__status on orders(status)`,
}
