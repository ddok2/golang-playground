// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - wallet.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

type bitcoin int

type Wallet struct {
	balance bitcoin
}

func (w *Wallet) Deposit(amount bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() bitcoin {
	return w.balance
}
