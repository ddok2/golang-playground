// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - wallet_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(bitcoin(10))
		assertBalance(t, wallet, bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		startingBalance := bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(bitcoin(10))

		assertBalance(t, wallet, bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err)
	})
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one..")
	}
}

func assertBalance(t *testing.T, wallet Wallet, want bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one..")
	}
}
