// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - test1.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "fmt"

// 6.1 배열

func main() {
	var x []float64

	x = make([]float64, 5, 10)

	fmt.Println(x)
}
