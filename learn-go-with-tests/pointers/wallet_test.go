// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - wallet_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"fmt"
	"testing"
)

// func TestWallet(t *testing.T) {
// 	wallet := Wallet{}
// 	wallet.Deposit(bitcoin(10))
//
// 	got := wallet.Balance()
// 	want := bitcoin(10)
//
// 	if got != want {
// 		t.Errorf("got %d wnat %d", got, want)
// 	}
// }

func TestResult(t *testing.T) {
	result := Result{UUID: "abc"}

	got := result.Code
	want := ""

	if got != want {
		t.Errorf("got %s wnat %s", got, want)
	}
	if got == want {
		fmt.Printf("same - got %s want %s", got, want)
		fmt.Println(result)
	}
}
