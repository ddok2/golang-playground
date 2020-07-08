// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - wallet.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
