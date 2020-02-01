// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - test1.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "fmt"

func main() {
	x := "안녕!"

	for i := 0; i < len(x); i++ {
		x := x[i]

		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // ÌuhËeu
		}
	}
}
