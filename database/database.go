package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewTable(db *sql.DB, createQuery string) {
	if db == nil {
		panic("DB is nil")
	}

	if err := execQuery(db, createQuery); err != nil {
		panic("Failed to create table: " + err.Error())
	}
}

func GetDB(dbName string) *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic("Failed to create DB: " + err.Error())
	}

	err = execQuery(db, "PRAGMA foreign_keys = ON")
	if err != nil {
		panic("Failed to create DB: " + err.Error())
	}

	return db
}

func execQuery(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
