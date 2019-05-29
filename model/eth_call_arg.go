package model

import "github.com/ethereum/go-ethereum/common"

// 以太坊 eth_call 的参数结构体
type CallArg struct {
	// common.Address 是以太坊依赖包的地址类型，其原型是 [20]byte 数组
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gas_price"`
	Value    string `json:"value"`
	Data     string `json:"data"`		// 这个就是 data
	Nonce    string `json:"nonce"`
}



