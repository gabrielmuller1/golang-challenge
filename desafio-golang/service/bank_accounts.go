package service

import (
	"github.com/eliasfeijo/desafio-golang-imersao/model"
	"github.com/eliasfeijo/desafio-golang-imersao/repository"
)

type BankAccounts interface {
	CreateBankAccount(number string) (*model.BankAccount, error)
}

type bankAccounts struct {
	repository repository.BankAccountsRepository
}

func NewBankAccounts(repository repository.BankAccountsRepository) BankAccounts {
	return &bankAccounts{repository: repository}
}

func (b bankAccounts) CreateBankAccount(number string) (*model.BankAccount, error) {
	id, err := b.repository.CreateBankAccount(number)
	if err != nil {
		return nil, err
	}
	bankAccount, err := b.repository.FindBankAccountById(id)
	if err != nil {
		return nil, err
	}
	return bankAccount, nil
}
