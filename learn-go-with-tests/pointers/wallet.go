// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - wallet.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "fmt"

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
