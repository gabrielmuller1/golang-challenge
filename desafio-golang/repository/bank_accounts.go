package repository

import (
	"database/sql"
	"log"

	"github.com/eliasfeijo/desafio-golang-imersao/database"
	"github.com/eliasfeijo/desafio-golang-imersao/model"
)

type BankAccountsRepository interface {
	CreateBankAccount(number string) (int64, error)
	FindBankAccountById(id int64) (*model.BankAccount, error)
	FindBankAccountByNumber(number string) (*model.BankAccount, error)
}

type bankAccountsRepository struct {
	db *sql.DB
}

var instanceBankAccounts BankAccountsRepository

func NewBankAccounts() BankAccountsRepository {
	if instanceBankAccounts == nil {
		db := database.GetConn()
		i := &bankAccountsRepository{db: db}
		instanceBankAccounts = i
	}
	return instanceBankAccounts
}

func (repository *bankAccountsRepository) CreateBankAccount(number string) (int64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO bank_accounts (number) VALUES (?)")
	if err != nil {
		log.Fatalf("Error preparing insert query: %v", err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(number)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return id, nil
}

func (repository *bankAccountsRepository) FindBankAccountById(id int64) (*model.BankAccount, error) {
	stmt, err := repository.db.Prepare("SELECT id, number, created_at FROM bank_accounts WHERE id = ?")
	if err != nil {
		log.Fatalf("Error preparing select query: %v", err)
		return nil, err
	}
	defer stmt.Close()

	bankAccount := model.BankAccount{}
	err = stmt.QueryRow(id).Scan(&bankAccount.ID, &bankAccount.Number, &bankAccount.CreatedAt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &bankAccount, nil
}

func (repository *bankAccountsRepository) FindBankAccountByNumber(number string) (*model.BankAccount, error) {
	stmt, err := repository.db.Prepare("SELECT id, number, created_at FROM bank_accounts WHERE number = ?")
	if err != nil {
		log.Fatalf("Error preparing select query: %v", err)
		return nil, err
	}
	defer stmt.Close()

	bankAccount := model.BankAccount{}
	err = stmt.QueryRow(number).Scan(&bankAccount.ID, &bankAccount.Number, &bankAccount.CreatedAt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &bankAccount, nil
}
