package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func InitTableIfNotExist(db *sqlx.DB) {

	schema := `
CREATE TABLE IF NOT EXISTS accounts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	balance REAL NOT NULL DEFAULT 0,
	is_active BOOLEAN DEFAULT 1
	
);

CREATE TABLE IF NOT EXISTS categories (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL UNIQUE,
	type TEXT NOT NULL,
	is_deleted BOOLEAN DEFAULT 0
);

CREATE TABLE IF NOT EXISTS transactions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	account_id INTEGER,
	category_id INTEGER,
	amount REAL NOT NULL,
	note TEXT,
	date TEXT NOT NULL,
	is_deleted BOOLEAN DEFAULT 0,
	FOREIGN KEY(account_id) REFERENCES accounts(id),
	FOREIGN KEY(category_id) REFERENCES categories(id)
);
`
	db.MustExec(schema)
	fmt.Println("Successfully created tables")
}
