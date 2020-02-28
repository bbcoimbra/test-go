package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Account struct {
	AccountID      int64  `json:"id,omitempty"`
	DocumentNumber string `json:"document_number"`
}

var accounts = make([]Account, 0)

var cur_account_id int64 = 0

func get_next_account_id() int64 {
	cur_account_id = cur_account_id + 1
	return cur_account_id
}

func accountCreationPath(w http.ResponseWriter, r *http.Request) {
	var newAccount Account
	var newID int64
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Persist Data
	newID = get_next_account_id()
	newAccount.AccountID = newID
	json.Unmarshal(reqBody, &newAccount)
	accounts = append(accounts, newAccount)

	fmt.Println(accounts)
	// Build response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAccount)
}

func accountGetItemPath(w http.ResponseWriter, r *http.Request) {
	account_id_str := mux.Vars(r)["accountId"]

	account_id, err := strconv.ParseInt(account_id_str, 10, 64)
	if err != nil {
		panic(err)
	}

	for _, account := range accounts {
		if account.AccountID == account_id {
			json.NewEncoder(w).Encode(account)
		}
	}
}
