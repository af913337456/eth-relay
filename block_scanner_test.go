package main

import (
	"eth-relay/dao"
	"testing"
	"fmt"
	"math/big"
)

/**
  作者(Author): 林冠宏 / 指尖下的幽灵
  Created on : 2018/12/11
*/

func TestBlockScanner_Start(t *testing.T) {
	requester := NewETHRPCRequester("https://mainnet.infura.io/v3/2e6d9331f74d472a9d47fe99f697ca2b")
	option := dao.MysqlOptions{
		Hostname:           "127.0.0.1",
		Port:               "3306",
		DbName:             "eth_relay",
		User:               "root",
		Password:           "123aaa",
		TablePrefix:        "eth_",
		MaxOpenConnections: 10,
		MaxIdleConnections: 5,
		ConnMaxLifetime:    15,
	}
	tables := []interface{}{}
	tables = append(tables, dao.Block{}, dao.Transaction{})
	mysql := dao.NewMqSQLConnector(&option, tables)
	scanner := NewBlockScanner(*requester, mysql)
	err := scanner.Start()
	if err != nil {
		panic(err)
	}
	select {}
}

func TestName2(t *testing.T) {
	fmt.Println(new(big.Int).SetString("68d5bc",16))
}

type Obj struct {
	T dao.Block
}

func TestPoint(t *testing.T) {
	o := &Obj{}
	o.T = dao.Block{BlockNumber:"456"}
	fmt.Println(o.T.BlockNumber)
	new := &dao.Block{BlockNumber:"123"}
	o.a(new)
	fmt.Println(o.T.BlockNumber)
}

func (o *Obj) a(block *dao.Block) {
	o.T = *block
}



