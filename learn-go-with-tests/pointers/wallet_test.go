// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - wallet_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(bitcoin(10))

	got := wallet.Balance()
	want := bitcoin(10)

	if got != want {
		t.Errorf("got %d wnat %d", got, want)
	}
}
