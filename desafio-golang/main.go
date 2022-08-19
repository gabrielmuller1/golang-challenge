package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eliasfeijo/desafio-golang-imersao/database"
	"github.com/eliasfeijo/desafio-golang-imersao/routes"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	cmd := os.Getenv("CMD")
	if cmd == "migrate" {
		err = database.Migrate()
		if err != nil {
			log.Fatalf("Error migrating database: %v", err)
			os.Exit(1)
		}
		log.Println("Database migrated successfully!")
		os.Exit(0)
	}

	router := mux.NewRouter()
	routes.SetupRoutesBankAccounts(router)
	routes.SetupRoutesTransfers(router)

	log.Println("API is running")
	http.ListenAndServe(":8000", router)
}
