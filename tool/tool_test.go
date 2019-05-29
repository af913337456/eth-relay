package tool

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"testing"
)

// 5
func Test_UnlockETHWallet(t *testing.T) {
	address := "0x590c3d81b70ddff32f74e51f14805915a4c0e2ed"
	keysDir := "../keystores"
	// 第一次演示密码错误的情况
	err1 := UnlockETHWallet(keysDir,address, "789")
	if err1 != nil {
		fmt.Println("第一次解锁错误：", err1.Error())
	} else {
		fmt.Println("第一次解锁成功!")
	}
	// 第二次密码正确，解锁成功
	err2 := UnlockETHWallet(keysDir,address, "13456aa")
	if err2 != nil {
		fmt.Println("第二次解锁错误：", err1.Error())
	} else {
		fmt.Println("第二次解锁成功!")
	}
	// 下面是签名的测试
	tx := types.NewTransaction( // 创建一个测试用的交易数据结构体
		123,                       // nonce 交易序列号
		common.Address{},          // to 接收者地址
		new(big.Int).SetInt64(10), // value 数值
		1000, // gasLimit
		new(big.Int).SetInt64(20), // gasPrice
		[]byte("交易"))              // data
	signTx, err := SignETHTransaction(address, tx)
	if err != nil {
		fmt.Println("签名失败!", err.Error())
		return
	}
	data, _ := json.Marshal(signTx)
	fmt.Println("签名成功\n", string(data))
}

func TestName(t *testing.T) {

}

func a(uint2 uint) {

}
