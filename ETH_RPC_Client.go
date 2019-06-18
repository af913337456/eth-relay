package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
)

type ETHRPCClient struct {
	NodeUrl string       // 代表节点的url链接
	client  *rpc.Client  // 代表 rpc 客户端实例
}

// NewETHRPCClient 代表的是新建一个 “RPC” 客户端
// 入参 nodeUrl，就是节点的链接，返回的是带有 * 号的ETHRPCClient指针对象
func NewETHRPCClient(nodeUrl string) *ETHRPCClient {
	// & 符号代表的是取指针
	client := &ETHRPCClient{
		NodeUrl:nodeUrl,
	}
	client.initRpc()  // 进行初始化 rpc 客户端句柄实体
	return client
}

// 初始化 rpc 请求实例
func (erc *ETHRPCClient) initRpc() {
	// 使用 go-ethereum 库中的 rpc 库来初始化
	// DialHTTP 的意思是使用 http 版本的 rpc 实现方式
	rpcClient,err := rpc.DialHTTP(erc.NodeUrl)
	if err != nil {
		// 初始化失败，终结程序，并将错误信息显示到控制台上面
		errInfo := fmt.Errorf("初始化 rpc client 失败%s",err.Error()).Error()
		panic(errInfo)
	}
	// 初始化成功，将新实例化的 rpc 句柄赋值给我们 ETHRPCClient 结构体里面的
	erc.client = rpcClient
}

// Go 语言语法中，大写字母开头的变量或者方法才能够被外部引用
// 小写字母的变量或方法只能内部调用
// GetRpc 方法是为了方便外部能够获取 client  *rpc.Client 来进行访问
func (erc *ETHRPCClient) GetRpc() *rpc.Client {
	if erc.client == nil {
		erc.initRpc()
	}
	return erc.client
}













