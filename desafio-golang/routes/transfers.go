package routes

import (
	"github.com/eliasfeijo/desafio-golang-imersao/controller"
	"github.com/eliasfeijo/desafio-golang-imersao/repository"
	"github.com/eliasfeijo/desafio-golang-imersao/service"
	"github.com/gorilla/mux"
)

func SetupRoutesTransfers(router *mux.Router) {
	transfersRepository := repository.NewTransfers()
	bankAccountsRepository := repository.NewBankAccounts()
	service := service.NewTransfers(transfersRepository, bankAccountsRepository)
	controller := controller.NewTransfers(service)

	router.HandleFunc("/bank-accounts/transfer", controller.CreateTransfer).Methods("POST")
}
