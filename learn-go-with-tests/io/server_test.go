// Copyright 2021. SUNGYUB.COM. All Rights Reserved.
//
// - server_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"reflect"
	"testing"
)

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
