// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - wallet_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(bitcoin(10))
		assertBalance(t, wallet, bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: bitcoin(20)}
		wallet.Withdraw(bitcoin(10))
		assertBalance(t, wallet, bitcoin(10))
	})
}
