package main

import (
	"encoding/json"
	"eth-relay/tool"
	"fmt"
	"math/big"
	"testing"
)

func TestNewETHRPCClient(t *testing.T) {
	// 首先是一个格式正确的链接测试初始化
	client2 := NewETHRPCClient("www.nihao.com").GetRpc()
	if client2 == nil {
		fmt.Println("初始化失败")
	}
	// 再次是 123://456 非法链接测试初始化
	client := NewETHRPCClient("123://456").GetRpc()
	if client == nil {
		fmt.Println("初始化失败")
	}
}

func Test_GetTransactionByHash(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/2e6d9331f74d472a9d47fe99f697ca2b"
	txHash := "0x53c5b03e392d6aa68a0df26b6d466ae8fbd1c2c5b74f9baae05434ec9a18a282"
	if txHash == "" || len(txHash) != 66 {
		// 这里演示，在调用 rpc 接口函数的时候，都要先进行入参的合法性判断
		fmt.Println("非法交易 hash 值")
		return
	}
	txInfo, err := NewETHRPCRequester(nodeUrl).GetTransactionByHash(txHash)
	if err != nil {
		// 查询失败，打印出信息
		fmt.Println("查询交易失败，信息是：", err.Error())
		return
	}
	// 查询成功，将 transaction 结果的结构体以 json 格式序列化，再以 string 格式输出
	json, _ := json.Marshal(txInfo)
	fmt.Println(string(json))
}

func Test_GetTransactions(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/2e6d9331f74d472a9d47fe99f697ca2b"
	txHash_1 := "0x53c5b03e392d6aa68a0df26b6d466ae8fbd1c2c5b74f9baae05434ec9a18a282"
	txHash_2 := "0x53c5b03e392d6aa68a0df26b6d466ae8fbd1c2c5b74f9baae05434ec9a18a281"
	txHash_3 := "0x711ddd5f223f970aa0ebc32304a880a8c2ec45ee134b4f41dd4da264f72e1afc"

	// txHash_1 是存在的，_2 是伪造的，_3 也是存在的
	txHashs := []string{}
	txHashs = append(txHashs, txHash_1, txHash_2, txHash_3)

	if txHashs == nil || len(txHashs) == 0 {
		// 这里演示，在调用 rpc 接口函数的时候，都要先进行入参的合法性判断
		fmt.Println("非法交易 hash 数组")
		return
	}
	txInfos, err := NewETHRPCRequester(nodeUrl).GetTransactions(txHashs)
	if err != nil {
		// 查询失败，打印出信息
		fmt.Println("查询交易失败，信息是：", err.Error())
		return
	}
	// 查询成功，将 transaction 结果的结构体以 json 格式序列化，再以 string 格式输出
	json, _ := json.Marshal(txInfos)
	fmt.Println(string(json))
}

// 单条单元测试函数
func Test_GetETHBalance(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/2e6d9331f74d472a9d47fe99f697ca2b"
	address := "0x0D0707963952f2fBA59dD06f2b425ace40b492Fe"
	if address == "" || len(address) != 42 {
		// 这里演示，在调用 rpc 接口函数的时候，都要先进行入参的合法性判断
		fmt.Println("非法交易 address 值")
		return
	}
	balance, err := NewETHRPCRequester(nodeUrl).GetETHBalance(address)
	if err != nil {
		// 查询失败，打印出信息
		fmt.Println("查询 eth 余额失败，信息是：", err.Error())
		return
	}
	fmt.Println(balance)
}

// 批量单元测试函数
func Test_GetETHBalances(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/2e6d9331f74d472a9d47fe99f697ca2b"

	address1 := "0x0D0707963952f2fBA59dD06f2b425ace40b492Fe" // 第一个地址
	address2 := "0xf89260db97765A00a343aba8e5682715804769ca" // 第二个地址

	address := []string{address1, address2}

	balance, err := NewETHRPCRequester(nodeUrl).GetETHBalances(address)
	if err != nil {
		// 查询失败，打印出信息
		fmt.Println("查询 eth 余额失败，信息是：", err.Error())
		return
	}
	fmt.Println(balance)
}

// 单元测试：批量获取代币值
func Test_GetERC20Balances(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/2e6d9331f74d472a9d47fe99f697ca2b"

	address := "0xc58AD8Ff428c354bb849d1dCf1EDCcAC3F102C8E"   // 钱包地址
	contract1 := "0x78021ABD9b06f0456cB9DB95a846C302c34f8b8D" // 合约地址1
	contract2 := "0xB8c77482e45F1F44dE1745F52C74426C631bDD52" // 合约地址2

	params := []ERC20BalanceRpcReq{}
	item := ERC20BalanceRpcReq{}
	item.ContractAddress = contract1
	item.UserAddress = address
	item.ContractDecimal = 18

	params = append(params, item)

	item.ContractAddress = contract2
	params = append(params, item)

	balance, err := NewETHRPCRequester(nodeUrl).GetERC20Balances(params)
	if err != nil {
		// 查询失败，打印出信息
		fmt.Println("查询 eth 余额失败，信息是：", err.Error())
		return
	}
	fmt.Println(balance)
}

// 单元测试：创建以太坊钱包
func Test_CreateETHWallet(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/2e6d9331f74d472a9d47fe99f697ca2b"
	address1, err := NewETHRPCRequester(nodeUrl).CreateETHWallet("13456") // 演示密码太短的错误
	if err != nil {
		fmt.Println("第一次，创建钱包失败", err.Error())
	} else {
		fmt.Println("第一次，创建钱包成功，以太坊地址是：", address1)
	}

	address2, err := NewETHRPCRequester(nodeUrl).CreateETHWallet("13456aa") // 创建成功
	if err != nil {
		fmt.Println("第二次，创建钱包失败", err.Error())
	} else {
		fmt.Println("第二次，创建钱包成功，以太坊地址是：", address2)
	}
}

// 单元测试： 获取 nonce
func Test_GetNonce(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/2e6d9331f74d472a9d47fe99f697ca2b"
	address := "0x0D0707963952f2fBA59dD06f2b425ace40b492Fe"
	if address == "" || len(address) != 42 {
		// 这里演示，在调用 rpc 接口函数的时候，都要先进行入参的合法性判断
		fmt.Println("非法交易 address 值")
		return
	}
	nonce, err := NewETHRPCRequester(nodeUrl).GetNonce(address)
	if err != nil {
		// 查询失败，打印出信息
		fmt.Println("查询 nonce 失败，信息是：", err.Error())
		return
	}
	fmt.Println(nonce)
}

// 单元测试： 转账 ETH
func Test_SendETHTransaction(t *testing.T) {
	nodeUrl := "https://ropsten.infura.io/v3/2e6d9331f74d472a9d47fe99f697ca2b" // ropsten 测试网络的节点链接
	from := "0x27d2ecd2e14e52243b68fcf2321f7a9550bdc0f2"                       // 这个地址就是当初获取测试代币的地址
	if from == "" || len(from) != 42 {
		// 这里演示，在调用 rpc 接口函数的时候，都要先进行入参的合法性判断
		fmt.Println("非法交易 address 值")
		return
	}
	to := "0xd8CCEFDac5F30f06C62ed13383e9563C482630Bc"
	value := "0.2" // 发送 0.2 个 ETH
	gasLimit := uint64(100000)
	gasPrice := uint64(36000000000)
	// 当前这笔交易消耗的油费最大值是 (gasLimit * gasPrice) / 10^18 ETH
	err := tool.UnlockETHWallet("./keystores", from, "123aaaaa") // 解锁钱包
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 下面发起交易转账
	txHash, err := NewETHRPCRequester(nodeUrl).SendETHTransaction(from, to, value, gasLimit, gasPrice)
	if err != nil {
		// 转账失败，打印出信息
		fmt.Println("ETH 转账失败，信息是：", err.Error())
		return
	}
	fmt.Println(txHash) // 打印出当前交易的 hash
}

// 单元测试：获取以太坊最新生成区块的区块号
func TestGetLatestBlockNumber(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/2e6d9331f74d472a9d47fe99f697ca2b"
	number, err := NewETHRPCRequester(nodeUrl).GetLatestBlockNumber()
	if err != nil {
		// 查询失败，打印出信息
		fmt.Println("获取区块号失败，信息是：", err.Error())
		return
	}
	fmt.Println("10进制: ", number.String())
}

func TestGetFullBlockInfo(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/2e6d9331f74d472a9d47fe99f697ca2b"

	requester := NewETHRPCRequester(nodeUrl)

	// 获取区块号
	number, _ := requester.GetLatestBlockNumber()

	// 获取区块信息
	fullBlock, err := requester.GetBlockInfoByNumber(number)

	if err != nil {
		// 查询失败，打印出信息
		fmt.Println("获取区块信息失败，信息是：", err.Error())
		return
	}
	// 查询成功，将 区块 结果的结构体以 json 格式序列化，再以 string 格式输出
	json1, _ := json.Marshal(fullBlock)
	fmt.Println("根据区块号获取区块信息", string(json1))

	// 根据区块 hash 获取区块信息
	fullBlock, err = requester.GetBlockInfoByHash(fullBlock.ParentHash)
	json2, _ := json.Marshal(fullBlock)
	fmt.Println("根据区块hash获取区块信息", string(json2))
}

func TestName(t *testing.T) {
	fmt.Println(fmt.Sprintf("%#x", new(big.Int).SetUint64(100)))
}

func TestLeeCode(t *testing.T) {
	fmt.Println(twoSum([]int{11, 7, 2, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	if nums == nil {
		return nil
	}
	size := len(nums) + 1
	nMap := map[int]int{}
	for i := 1; i < size; i++ {
		nMap[nums[i-1]] = i
	}
	for i := 1; i < size; i++ {
		index := nMap[target-nums[i-1]]
		if index != 0 {
			return []int{i - 1, index - 1}
		}
	}
	return nil
}
