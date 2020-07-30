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
		t.Run("unknown word", func(t *testing.T) {
			_, got := dictionary.Search("unknown")

			assertError(t, got, ErrNotFound)
		})
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
 go test
# _/Users/sung/Development/01.sung/golang-playground/learn-go-with-tests/maps [_/Users/sung/Development/01.sung/golang-playground/learn-go-with-tests/maps.test]
./dictionary_test.go:35:24: dictionary.Add(word, definition) used as value
./dictionary_test.go:45:24: dictionary.Add(word, "new test") used as value
./dictionary_test.go:47:23: undefined: ErrWordExists
FAIL    _/Users/sung/Development/01.sung/golang-playground/learn-go-with-tests/maps [build failed]
*/
