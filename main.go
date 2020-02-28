package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
)

const PAYMENT_TYPE = 4

type Account struct {
	AccountID      int64  `json:"id,omitempty"`
	DocumentNumber string `json:"document_number"`
}

type OperationType struct {
	OperationTypeID int64  `json:"account_id,omitempty"`
	Description     string `json:"description"`
}

type Transaction struct {
	TransactionID   int64           `json:"transaction_id,omitempty"`
	AccountID       int64           `json:"account_id"`
	OperationTypeID int64           `json:"operation_type_id"`
	Amount          float64         `json:"amount"`
	EventDate       time.Time       `json:"event_date,omitempty"`
	DecimalAmount   decimal.Decimal `json:"decimal_amount,omitempty"`
}

var accounts = make([]Account, 0)

var cur_account_id int64 = 0

var cur_transaction_id int64 = 0

var transactions = make([]Transaction, 0)

var operation_types = []OperationType{
	{
		OperationTypeID: 1,
		Description:     "COMPRA A VISTA",
	},
	{
		OperationTypeID: 2,
		Description:     "COMPRA PARCELADA",
	},
	{
		OperationTypeID: 3,
		Description:     "SAQUE",
	},
	{
		OperationTypeID: 4,
		Description:     "PAGAMENTO",
	},
}

func homePath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Pismo test")
}

func get_next_account_id() int64 {
	cur_account_id = cur_account_id + 1
	return cur_account_id
}

func get_next_transaction_id() int64 {
	cur_transaction_id = cur_transaction_id + 1
	return cur_transaction_id
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

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePath)
	router.HandleFunc("/accounts", accountCreationPath).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", accountGetItemPath).Methods("GET")
	router.HandleFunc("/transactions", transactionCreationPath).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
