// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - play_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package hello

import "testing"

func TestTest2(t *testing.T) {
	want := [...]rune{
		'H', 'e', 'l', 'l', 'o', ',', ' ', '썽', '썽', '이',
	}

	index := 1
	if got := Test2(); got[index] != want[index] {
		t.Errorf("Test2()[%q] = %q, want %q",
			index, got, want)
	}
}

/*
 go test
['H']
['H' 'e']
['H' 'e' 'l']
['H' 'e' 'l' 'l']
['H' 'e' 'l' 'l' 'o']
['H' 'e' 'l' 'l' 'o' ',']
['H' 'e' 'l' 'l' 'o' ',' ' ']
['H' 'e' 'l' 'l' 'o' ',' ' ' '썽']
['H' 'e' 'l' 'l' 'o' ',' ' ' '썽' '썽']
['H' 'e' 'l' 'l' 'o' ',' ' ' '썽' '썽' '이']
PASS
ok      hello   0.386s

*/
