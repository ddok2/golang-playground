// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - hello_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got: %q \nwant: %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sung")
		want := "Hello, Sung"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empth string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}
