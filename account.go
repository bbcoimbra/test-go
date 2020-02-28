package main

import (
	"encoding/json"
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

func persistAccount(reqBody []byte) *Account {
	var newAccount Account
	var newID int64

	newID = get_next_account_id()
	newAccount.AccountID = newID
	json.Unmarshal(reqBody, &newAccount)
	accounts = append(accounts, newAccount)

	return &newAccount
}

func findAccount(accountId int64) *Account {
	var accountToReturn *Account
	for _, account := range accounts {
		if account.AccountID == accountId {
			accountToReturn = &account
		}
	}

	return accountToReturn
}
