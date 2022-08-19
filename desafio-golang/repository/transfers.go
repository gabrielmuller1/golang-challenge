package repository

import (
	"database/sql"
	"log"

	"github.com/eliasfeijo/desafio-golang-imersao/database"
)

type TransfersRepository interface {
	CreateTransfer(from int64, to int64, amount float64) error
	Balance(bankAccountId int64) (float64, error)
}

type transfersRepository struct {
	db *sql.DB
}

var instanceTransfers TransfersRepository

func NewTransfers() TransfersRepository {
	if instanceTransfers == nil {
		db := database.GetConn()
		i := &transfersRepository{db: db}
		instanceTransfers = i
	}
	return instanceTransfers
}

func (repository *transfersRepository) CreateTransfer(fromId int64, toId int64, amount float64) error {
	stmt, err := repository.db.Prepare("INSERT INTO transfers (from_id, to_id, amount) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalf("Error preparing insert query: %v", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(fromId, toId, amount)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (repository *transfersRepository) Balance(bankAccountId int64) (float64, error) {
	stmt, err := repository.db.Prepare("SELECT SUM(amount) FROM transfers WHERE from_id = ?")
	if err != nil {
		log.Fatalf("Error preparing select query: %v", err)
		return 0, err
	}
	defer stmt.Close()

	var paidAmount float64
	err = stmt.QueryRow(bankAccountId).Scan(&paidAmount)
	if err != nil {
		paidAmount = 0
	}

	stmt, err = repository.db.Prepare("SELECT SUM(amount) FROM transfers WHERE to_id = ?")
	if err != nil {
		log.Fatalf("Error preparing select query: %v", err)
		return 0, err
	}
	defer stmt.Close()

	var receivedAmount float64
	err = stmt.QueryRow(bankAccountId).Scan(&receivedAmount)
	if err != nil {
		receivedAmount = 0
	}

	balance := receivedAmount - paidAmount

	return balance, nil
}
