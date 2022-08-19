package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/eliasfeijo/desafio-golang-imersao/service"
)

type BankAccounts interface {
	CreateBankAccount(w http.ResponseWriter, r *http.Request)
}

type bankAccounts struct {
	service service.BankAccounts
}

func NewBankAccounts(service service.BankAccounts) BankAccounts {
	return &bankAccounts{service: service}
}

type CreateBankAccountPayload struct {
	Number string `json:"account_number"`
}

func (b bankAccounts) CreateBankAccount(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")

	var payload CreateBankAccountPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if payload.Number == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("'account_number' param is required")
		return
	}

	bankAccount, err := b.service.CreateBankAccount(payload.Number)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Error creating BankAccount")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bankAccount)
}
