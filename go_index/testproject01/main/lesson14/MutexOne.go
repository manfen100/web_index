package main

import (
	"fmt"
	"sync"
)

type BankV2 struct {
	balance int
	m       sync.Mutex
}

func (b *BankV2) Deposit(amount int) {
	b.m.Lock()
	b.balance += amount
	b.m.Unlock()
}

func (b *BankV2) Balance() int {
	return b.balance
}
func main() {
	var wg sync.WaitGroup
	b := &BankV2{}

	n := 1000
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			b.Deposit(1000)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(b.Balance())
}
