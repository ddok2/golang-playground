// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - raceConditions.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"fmt"
	"time"
)

func main() {
	var data int

	go func() {
		data++
	}()

	time.Sleep(1 * time.Second)

	if data == 0 {
		fmt.Printf("the value is %v. \n", data)
	}
}

// line 20은 잘못되었다.
