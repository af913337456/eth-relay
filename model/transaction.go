package model

type Transaction struct {
	Hash             string    	`json:"hash"`
	Nonce            string 	`json:"nonce"`
	BlockHash        string    	`json:"blockHash"`
	BlockNumber      string 	`json:"blockNumber"`
	TransactionIndex string 	`json:"transactionIndex"`
	From             string    	`json:"from"`
	To               string    	`json:"to"`
	Value            string 	`json:"value"`
	GasPrice         string 	`json:"gasPrice"`
	Gas              string 	`json:"gas"`
	Input            string    	`json:"input"`
}























