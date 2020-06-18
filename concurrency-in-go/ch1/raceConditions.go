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
	} else {
		fmt.Printf("the value is %v. \n", data)
	}
}

// line 20은 잘못되었다.
// data 변수를 증가시키는 goroutine
// data 값이 0인지 확인하는 if 구문
// 출력할 data의 값을 가져오는 fmt.Printf 구문
