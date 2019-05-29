package main

import (
	"math/big"
	"sync"
)

// 管理器结构体
type NonceManager struct {
	// lock 是互斥锁，go 的 map 类型不是线程安全的，
	// 在读写 map 的时候，我们要考虑上多协程并发的情况
	lock sync.Mutex
	// 采用整形大数来存储 nonce
	nonceMemCache map[string]*big.Int
}

func NewNonceManager() *NonceManager {
	return &NonceManager{
		lock: sync.Mutex{}, // 实例化互斥锁
	}
}
//  设置 nonce
func (n *NonceManager) SetNonce(address string,nonce *big.Int) {
	if n.nonceMemCache == nil {
		n.nonceMemCache = map[string]*big.Int{}
	}
	n.lock.Lock()         // 加锁
	defer n.lock.Unlock() // 当该函数执行完毕，进行解锁
	n.nonceMemCache[address] = nonce
}

// 根据以太坊地址获取 nonce
func (n *NonceManager) GetNonce(address string) *big.Int {
	if n.nonceMemCache == nil {
		n.nonceMemCache = map[string]*big.Int{}
	}
	n.lock.Lock()         // 加锁
	defer n.lock.Unlock() // 当该函数执行完毕，进行解锁
	return n.nonceMemCache[address]
}

// nonce 进行自增 1
func (n *NonceManager) PlusNonce(address string) {
	if n.nonceMemCache == nil {
		n.nonceMemCache = map[string]*big.Int{}
	}
	n.lock.Lock()         // 加锁
	defer n.lock.Unlock() // 当该函数执行完毕，进行解锁
	oldNonce := n.nonceMemCache[address]
	if oldNonce == nil {
		return
	}
	newNonce := oldNonce.Add(oldNonce, big.NewInt(int64(1)))
	n.nonceMemCache[address] = newNonce
}
