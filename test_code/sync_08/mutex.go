package main

import (
	"fmt"
	"sync"
	"time"
)

type Bank struct {
	sync.Mutex
	saving map[string]int
}

func NewBank() *Bank {
	return &Bank{
		saving: make(map[string]int),
	}
}

// 存款
func (b *Bank)Deposit(name string, amount int)  {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	if _, ok :=b.saving[name];!ok {
		b.saving[name] = 0
	}
	b.saving[name] += amount
}

// 取款
func (b *Bank)Withdraw(name string, amount int) int {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	if _, ok := b.saving[name];!ok{
		return 0
	}
	if b.saving[name]<amount {
		amount = b.saving[name]
	}
	b.saving[name] -= amount
	return amount
}

// 查询余额
func (b *Bank)Query(name string) int {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	if _, ok := b.saving[name];!ok{
		return 0
	}
	return b.saving[name]
}

func main()  {
	b := NewBank()
	go b.Deposit("xxx", 100)
	go b.Withdraw("xxx", 20)
	go b.Deposit("xxx", 20000)

	time.Sleep(time.Second)

	fmt.Println("xxx has : %d", b.Query("xxx"))
	fmt.Println("ddd has : %d", b.Query("ddd"))
}