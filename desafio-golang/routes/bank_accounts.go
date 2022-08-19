package routes

import (
	"github.com/eliasfeijo/desafio-golang-imersao/controller"
	"github.com/eliasfeijo/desafio-golang-imersao/repository"
	"github.com/eliasfeijo/desafio-golang-imersao/service"
	"github.com/gorilla/mux"
)

func SetupRoutesBankAccounts(router *mux.Router) {
	repository := repository.NewBankAccounts()
	service := service.NewBankAccounts(repository)
	controller := controller.NewBankAccounts(service)

	router.HandleFunc("/bank-accounts", controller.CreateBankAccount).Methods("POST")
}
