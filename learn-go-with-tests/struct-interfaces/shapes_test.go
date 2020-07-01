// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - shapes_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
