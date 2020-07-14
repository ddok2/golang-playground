// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - ch1-18.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "fmt"

var originalCount int
var eatenCount int

func main() {
	originalCount = 10
	eatenCount = 4

	fmt.Println("I stated with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apple left.")
}
