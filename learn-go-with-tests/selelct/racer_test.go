// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - racer_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package selelct

import "testing"

func TestRacer(t *testing.T) {
	url1 := "https://www.facebook.com"
	url2 := "https://www.google.com"

	want := url2
	got := Racer(url1, url2)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
