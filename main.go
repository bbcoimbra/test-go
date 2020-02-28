package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/accounts", accountCreationPath).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", accountGetItemPath).Methods("GET")
	router.HandleFunc("/transactions", transactionCreationPath).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
