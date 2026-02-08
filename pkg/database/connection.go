package database

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite", "fintr.db")
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	return db, nil
}
