// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - dictionary_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/**
 go test
# _/Users/sung/Development/01.sung/golang-playground/learn-go-with-tests/maps [_/Users/sung/Development/01.sung/golang-playground/learn-go-with-tests/maps.test]
./dictionary_test.go:14:10: assignment mismatch: 2 variables but dictionary.Search returns 1 values
./dictionary_test.go:21:10: assignment mismatch: 2 variables but dictionary.Search returns 1 values
FAIL    _/Users/sung/Development/01.sung/golang-playground/learn-go-with-tests/maps [build failed]
*/
