package service

import (
	"errors"
	"fmt"

	"github.com/eliasfeijo/desafio-golang-imersao/repository"
)

type Transfers interface {
	CreateTransfer(accountNumberFrom string, accountNumberTo string, amount float64) (float64, float64, error)
}

type transfers struct {
	transfersRepository    repository.TransfersRepository
	bankAccountsRepository repository.BankAccountsRepository
}

func NewTransfers(transfersRepository repository.TransfersRepository, bankAccountsRepository repository.BankAccountsRepository) Transfers {
	return &transfers{transfersRepository: transfersRepository, bankAccountsRepository: bankAccountsRepository}
}

func (t transfers) CreateTransfer(accountNumberFrom string, accountNumberTo string, amount float64) (float64, float64, error) {

	from, err := t.bankAccountsRepository.FindBankAccountByNumber(accountNumberFrom)
	if err != nil {
		return 0, 0, errors.New("service.transfers: Invalid 'from' account number")
	}

	to, err := t.bankAccountsRepository.FindBankAccountByNumber(accountNumberTo)
	if err != nil {
		return 0, 0, errors.New("service.transfers: Invalid 'to' account number")
	}

	err = t.transfersRepository.CreateTransfer(from.ID, to.ID, amount)
	if err != nil {
		return 0, 0, fmt.Errorf("service.transfers: Error creating Transfer: %v", err)
	}

	balanceFrom, err := t.transfersRepository.Balance(from.ID)
	if err != nil {
		return 0, 0, fmt.Errorf("service.transfers: Error calculating balance from: %v", err)
	}

	balanceTo, err := t.transfersRepository.Balance(to.ID)
	if err != nil {
		return 0, 0, fmt.Errorf("service.transfers: Error calculating balance to: %v", err)
	}

	return balanceFrom, balanceTo, nil
}
