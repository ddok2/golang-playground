// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - test1.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "fmt"

func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	}
}

func g(x int) int {
	return x + 0
}

func f() int {
	return 0
}
