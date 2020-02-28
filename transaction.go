package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

const PAYMENT_TYPE = 4

type Transaction struct {
	TransactionID   int64           `json:"transaction_id,omitempty"`
	AccountID       int64           `json:"account_id"`
	OperationTypeID int64           `json:"operation_type_id"`
	Amount          float64         `json:"amount"`
	EventDate       time.Time       `json:"event_date,omitempty"`
	DecimalAmount   decimal.Decimal `json:"decimal_amount,omitempty"`
}

var cur_transaction_id int64 = 0

var transactions = make([]Transaction, 0)

func get_next_transaction_id() int64 {
	cur_transaction_id = cur_transaction_id + 1
	return cur_transaction_id
}

func transactionCreationPath(w http.ResponseWriter, r *http.Request) {
	var newTransaction Transaction
	var newID int64
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Data transformations
	newID = get_next_transaction_id()
	newTransaction.TransactionID = newID
	json.Unmarshal(reqBody, &newTransaction)
	newTransaction.EventDate = time.Now()
	if newTransaction.OperationTypeID == PAYMENT_TYPE {
		newTransaction.DecimalAmount = decimal.NewFromFloat(newTransaction.Amount)
	} else {
		newTransaction.DecimalAmount = decimal.NewFromFloat(newTransaction.Amount).Neg()
	}

	// Persist Data
	transactions = append(transactions, newTransaction)

	fmt.Println(transactions)
	// Build response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTransaction)
}
