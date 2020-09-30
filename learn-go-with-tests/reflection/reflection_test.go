// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - reflection_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct{ Name string }{"Sung"},
			[]string{"Sung"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Sung", "Seoul"},
			[]string{"Sung", "Seoul"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Sung", 34},
			[]string{"Sung"},
		},
		{
			"Nested fields",
			Person{
				"Sung",
				Profile{34, "Seoul"},
			},
			[]string{"Sung", "Seoul"},
		},
		{
			"Pointers to things",
			&Person{
				"Sung",
				Profile{34, "Seoul"},
			},
			[]string{"Sung", "Seoul"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
