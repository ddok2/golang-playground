// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - play.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package hello

import "fmt"

// Test2
func Test2() []rune {
	var runes []rune

	for _, r := range "Hello, 썽썽이" {
		runes = append(runes, r)

		fmt.Printf("%q\n", runes)
	}

	return runes
}
