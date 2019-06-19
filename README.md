### 以太坊中继 (ETH-Relay)

书籍《区块链以太坊DApp开发实战》的 Demo 源码。

* 使用 Go 语言编写
* Go 版本为： 1.11
* MySQL 版本为：5.7.19，引擎选择 Innodb
* 代码开发工具是：Goland


### 启动

代码主要采用单元测试的方式运行。并没编写 `main` 函数，读者可以自行拓展。区块扫描部分需要依赖到 MySQL 数据库。

其中 `eth-relay/block_scanner_test.go` 的 `TestBlockScanner_Start` 函数是区块遍历入口函数

### 其它

欢迎大家在 issues 中给我提一些有用的建议 或 贡献代码，我们一起维护它。谢谢
