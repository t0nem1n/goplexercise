package exer9_1

import (
	"fmt"
	"sync"
)

var (
	deposits  = make(chan int)
	balances  = make(chan int)
	withdraws = make(chan withdrawInfo)
)

type withdrawInfo struct {
	amount  int
	success chan bool
}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	w := withdrawInfo{amount, make(chan bool)}
	withdraws <- w
	return <-w.success
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-withdraws:
			if balance >= w.amount {
				balance -= w.amount
				w.success <- true
			} else {
				w.success <- false
			}
		}
	}
}

func Bank() {
	go teller()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		Deposit(200)
		fmt.Println(" = ", Balance())
		wg.Done()
	}()

	go func() {
		Deposit(100)
		fmt.Println(" = ", Balance())
		fmt.Println(Withdraw(200))
		wg.Done()
	}()

	wg.Wait()
}
