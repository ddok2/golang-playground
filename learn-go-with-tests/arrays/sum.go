// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - sum.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
