// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - di_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sung")

	got := buffer.String()
	want := "Hello, Sung"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
