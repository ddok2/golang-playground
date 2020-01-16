// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - main.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "fmt"

const _boilingF = 212.0

func main() {
	var f = _boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)

	// 출력:
	// boiling point = 212°F or 100°C
}
