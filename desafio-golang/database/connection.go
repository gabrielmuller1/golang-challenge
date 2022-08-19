package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetConn() *sql.DB {
	return db
}

func ConnectDatabase() (*sql.DB, error) {
	dbConn, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return nil, err
	}

	db = dbConn

	return db, nil
}

func Migrate() error {
	b, err := os.ReadFile("./database/migration/01_create_bank_accounts.sql")
	if err != nil {
		return err
	}
	queryCreateBankAccounts := string(b)
	db.Exec(queryCreateBankAccounts)
	b, err = os.ReadFile("./database/migration/02_create_transfers.sql")
	if err != nil {
		return err
	}
	queryCreateTransfers := string(b)
	db.Exec(queryCreateTransfers)
	return nil
}
