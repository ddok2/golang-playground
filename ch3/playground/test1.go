// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - test1.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "fmt"

// 6.1 배열

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5, 6, 7, 8, 9)

	fmt.Println(slice1, slice2)
}
