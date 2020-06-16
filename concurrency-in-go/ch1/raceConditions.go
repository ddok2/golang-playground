// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - raceConditions.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "fmt"

func main() {
	var data int

	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v. \n", data)
	}
}

// 'the value is 0' is printed
// nothing is printed
// 'the value is 1' is printed
