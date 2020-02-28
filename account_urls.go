package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func accountCreationPath(w http.ResponseWriter, r *http.Request) {
	var newAccount *Account
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Persist Data
	newAccount = persistAccount(reqBody)
	fmt.Println(accounts) // TODO Remove debug code

	// Build response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(*newAccount)
}

func accountGetItemPath(w http.ResponseWriter, r *http.Request) {
	var account *Account
	account_id_str := mux.Vars(r)["accountId"]

	account_id, err := strconv.ParseInt(account_id_str, 10, 64)
	if err != nil {
		panic(err)
	}

	account = findAccount(account_id)
	json.NewEncoder(w).Encode(account)
}
