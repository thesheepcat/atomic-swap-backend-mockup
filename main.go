package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	restApiRequestsHandlers()
	fmt.Println("Server is up and running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Define a handler for each endpoint
func restApiRequestsHandlers() {
	http.HandleFunc("/is-online", isOnlineEndpoint)
	http.HandleFunc("/initiate-swap-contract", initiateSwapContractEnpoint)
}

// Check if BTC / KAS network are available
func isOnlineEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if isOnline() {
		json.NewEncoder(w).Encode(true)
	} else {
		json.NewEncoder(w).Encode(false)
	}
}

// Dummy function  - To be replaced
func isOnline() bool {
	return true
}

// Initiate swap contract
func initiateSwapContractEnpoint(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
	}
	var recipientAndAmountData recipientAddressAndAmount
	json.Unmarshal(reqBody, &recipientAndAmountData)
	contractTransactionData, err := initiateSwapContract(recipientAndAmountData.RecipientAddress, recipientAndAmountData.Amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contractTransactionData)
	}
}

// Dummy function  - To be replaced
func initiateSwapContract(recipientAddress string, amount float32) (initiateContractTransaction, error) {
	fmt.Println("Initiate contract requested: ", recipientAddress, amount)

	contractTransaction := initiateContractTransaction{
		Secret:              "3e0b064c97247732a3b345ce7b2a835d928623cb2871c26db4c2539a38e61a16",
		SecretHash:          "29c36b8dd380e0426bdc1d834e74a630bfd5d111",
		TransactionFee:      "0.0012345 BTC (0.00020000 BTC/kB)",
		Contract:            "63a61429c36b8dd380e0426bdc1d834e74a630bfd5d1118876a914eBcf822c4a2cdB5f6a6b9c4a59b74d66461da5816704d728bd59b17576a91406fb26221375b1cbe2c17c14f1bc2510b9f8f8ff6888ac",
		ContractTransaction: "34e74a630bfd5d1118876a914eBcf822c4a263a61429c36b8dd380e0426bdc1d834e74a630bfd5d1118876a914eBcf822c4a2cdB5f6a6b9c4a59b74d66461da5816704d728bd59b17576a91406fb2d5d1118876a914eBcf822c4a36b8dd380e0426bdc1d834e74a630bfd5d1118876a914eBcf822c4a2cdB5f6a6b9c4a59b74d66461da5816704d728bd59b17576a91406fb2d5d1118876a914eBcf822c4a6221375b1cbe2c17c14f1bc2510b6461da5816704d728bd",
	}
	// Save transaction on lastCreatedTransaction variable
	return contractTransaction, nil
}

type initiateContractTransaction struct {
	Secret              string `json:"Secret"`
	SecretHash          string `json:"SecretHash"`
	TransactionFee      string `json:"TransactionFee"`
	Contract            string `json:"Contract"`
	ContractTransaction string `json:"ContractTransaction"`
}

type recipientAddressAndAmount struct {
	RecipientAddress string  `json:"RecipientAddress"`
	Amount           float32 `json:"Amount"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////


type buildContractOutput struct {
  //if partecipate this will be an empty string
  Secret              *string `json:"Secret"`
  SecretHash          string  `json:"SecretHash"`
  Contract            string  `json:"Contract"`
  TxID                string  `json:"ContractTransactionID"`
  Tx                  string  `json:"ContractTransaction"`
  TxFee               string  `json:"TransactionFee"`
}

type buildContractInput struct {
  Them        string  `json:"RecipientAddress"`
  Amount      string  `json:"Amount"`
  //if nil or empty string I'll initiate a conctract
  //partecipate otherwise
  SecretHash  *string `json:"secretHash"`
}
type spendContractOutput struct {
  Tx  string `json:"SpendTransaction"`
  TxID  string `json:"SpendTransactionID"`
  TxFee    string `json:"TransactionFee"`
}
type spendContractInput struct {
  //if nil or empty string I'll start refund
  //redeem otherwise
  Secret              *string `json:"Secret"`
  Contract            string  `json:"Contract"`
  Tx string  `json:"ContractTransaction"`
}
type auditContractInput struct {
  Contract    string `json:"Contract"`
  Tx  string `json:"ContractTx"`
}
type  auditContractOutput struct {
  ContractAddress   string `json:"ContractAddress"`
  //if I don't konw the address this is empty string
  RecipientAddress  string `json:"RecipientAddress"`
  Recipient2b       string `json:"RecipientBlake2b"`
  Amount            string `json:"ContractAmount"`
  //if I don't konw the address this is empty string
  RefundAddress     string `json:"RefundAddress"`
  Refund2b          string `json:"RefundBlake2b"`
  SecretHash        string `json:"SecretHash"`
  LockTime          string `json:"LockTime"`
  TxId              string `json:"TxId"`
  DaaScore          string `json:"DaaScore"`
}
type extractSecretInput struct {
  Tx string `json:"Transaction"`
  SecretHash string `json:"SecretHash"`
}
type extractSecretOutput struct {
  Secret  string `json:"Secret"`
}
