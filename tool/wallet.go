package tool

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"strings"
	"math/big"
)

// 全局的保存了已经解锁成功了的钱包 map 集合变量
var ETHUnlockMap map[string]accounts.Account

// 全局的对应 keystore 实例
var UnlockKs *keystore.KeyStore

// 解锁以太坊钱包，传入钱包地址和对应的 keystore密码
func UnlockETHWallet(keysDir string,address, password string) error {
	if UnlockKs == nil {
		UnlockKs = keystore.NewKeyStore(
			// 服务端存储 keystore 文件的目录
			// 这些配置类的信息可以由配置文件指定
			keysDir,
			keystore.StandardScryptN,
			keystore.StandardScryptP)
		if UnlockKs == nil {
			return errors.New("ks is nil")
		}
	}
	unlock := accounts.Account{Address: common.HexToAddress(address)}
	// ks.Unlock 调用 keystore.go 的解锁函数，解锁出的私钥将存储在它里面的变量中
	if err := UnlockKs.Unlock(unlock, password); nil != err {
		return errors.New("unlock err ： " + err.Error())
	}
	if ETHUnlockMap == nil {
		ETHUnlockMap = map[string]accounts.Account{}
	}
	ETHUnlockMap[address] = unlock // 解锁成功，存储
	return nil
}

// 签名交易数据结构体 types.Transaction
func SignETHTransaction(address string, transaction *types.Transaction) (*types.Transaction, error) {
	if UnlockKs == nil {
		return nil, errors.New("you need to init keystore first!")
	}
	account := ETHUnlockMap[address]
	if !common.IsHexAddress(account.Address.String()) {
		// 判断当前地址钱包是否解锁了
		return nil, errors.New("account need to unlock first!")
	}
	return UnlockKs.SignTx(account, transaction, nil) // 调用签名函数
}

// 根据代币的 decimal 得出乘上 10^decimal 后的值
// value 是包含浮点数的。例如 0.5 个 ETH
func GetRealDecimalValue(value string,decimal int) string {
	if strings.Contains(value, ".") {
		// 小数
		arr := strings.Split(value, ".")
		if len(arr) != 2 {
			return ""
		}
		num := len(arr[1])
		left := decimal - num
		return arr[0] + arr[1] + strings.Repeat("0", left)
	} else {
		// 整数
		return value + strings.Repeat("0", decimal)
	}
}

// 构建符合“ERC20”标准的“transfer”合约函数的“data”入参
func BuildERC20TransferData(value,receiver string,decimal int) string {

	realValue := GetRealDecimalValue(value,decimal)  // 将 value 转为乘上 10^decimal 的格式
	valueBig, _ := new(big.Int).SetString(realValue, 10)

	// 按照 “交易参数的说明”小节中的讲解，进行构建
	methodId := "0xa9059cbb" // "0xa9059cbb" 是 transfer 的 methodId
	param1 := common.HexToHash(receiver).String()[2:]  // 第一个参数，收款者地址
	param2 := common.BytesToHash(valueBig.Bytes()).String()[2:] // 第二个参数，交易的数值
	return methodId + param1 + param2
}






