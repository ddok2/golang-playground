// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - deadlocks.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func main() {
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&a, &b)
	wg.Wait()
}

// 24번째 여기서는 입력값 v1 을 위해 임계 영역에 진입하려고 시도한다.
// 25번째 printSum 함수가 리턴되기전에 임계 영역을 벗어나도록 하기위해 defer 문을 사용하고 있다.
// 27번째 여기서 일정시간동안 sleep 하면서 임의의 작업을 수행하는 것처럼 시뮬레이션한다.
