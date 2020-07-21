// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - wallet.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "errors"

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

func (w *Wallet) Withdraw(amount bitcoin) error {
	if amount > w.balance {
		return errors.New("oh no..")
	}
	w.balance -= amount
	return nil
}
