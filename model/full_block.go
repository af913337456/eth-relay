package model

import "eth-relay/dao"

// 根据文档定义出区块信息的结构体
type FullBlock struct {
	Number           string   	`json:"number"`
	Hash             string 	`json:"hash"`
	ParentHash       string 	`json:"parentHash"`
	Nonce            string   	`json:"nonce"`
	Sha3Uncles       string   	`json:"sha3Uncles"`
	LogsBloom        string   	`json:"logsBloom"`
	TransactionsRoot string   	`json:"transactionsRoot"`
	ReceiptsRoot     string   	`json:"stateRoot"`
	Miner            string   	`json:"miner"`
	Difficulty       string   	`json:"difficulty"`
	TotalDifficulty  string   	`json:"totalDifficulty"`
	ExtraData        string   	`json:"extraData"`
	Size             string   	`json:"size"`
	GasLimit         string   	`json:"gasLimit"`
	GasUsed          string   	`json:"gasUsed"`
	Timestamp        string   	`json:"timestamp"`
	Uncles           []string   `json:"uncles"`
	Transactions 	 []dao.Transaction 	`json:"transactions"`
}