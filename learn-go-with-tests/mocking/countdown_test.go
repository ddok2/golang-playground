// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - countdown_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {

	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
