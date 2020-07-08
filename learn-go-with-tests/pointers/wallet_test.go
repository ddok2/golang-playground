// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - wallet_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	if got != want {
		t.Errorf("got %d wnat %d", got, want)
	}
}

/**
 go test
address of balance in Deposit is 0xc0000a6070
address of balance in test is 0xc0000a6070
PASS
*/
