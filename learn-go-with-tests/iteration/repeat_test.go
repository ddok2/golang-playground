// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - repeat_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package iteration

import "testing"

func TestRepaet(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}
