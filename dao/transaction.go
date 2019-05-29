package dao

type Transaction struct {
	Id               int64  `json:"id"`               // 主键
	Hash             string `json:"hash"`             // 交易的 hash
	Nonce            string `json:"nonce"`            // 交易的序列号
	BlockHash        string `json:"blockHash"`        // 当前交易被打包的区块的hash
	BlockNumber      string `json:"blockNumber"`      // 当前交易被打包在的区块的区块号
	TransactionIndex string `json:"transactionIndex"` // 当前交易在区块已打包交易数组中的下标
	From             string `json:"from"`             // 交易发起者地址
	To               string `json:"to"`               // 交易接收者地址
	Value            string `json:"value"`            // 交易的数值
	GasPrice         string `json:"gasPrice"`         // gasPrice
	Gas              string `json:"gas"`              // gasLimit
	Input            string `xorm:"text" json:"input"` // data
}
