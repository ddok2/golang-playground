// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - hello_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sung")
	want := "Hello, Sung"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
