// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - numeral_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package roman_numerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	got := ConvertToRoman(1)
	want := "I"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func ConvertToRoman(arabic int) string {
	return ""
}
