// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - sum.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	return
}