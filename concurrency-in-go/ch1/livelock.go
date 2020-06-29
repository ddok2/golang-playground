// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - livelock.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func test() {
	cadence := sync.NewCond(&sync.Mutex{})

	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()

		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!!")
			return true
		}
		takeStep()
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("left", &left, out)
	}
	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("right", &right, out)
	}

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()
		defer walking.Done()

		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	var pepleInHallway sync.WaitGroup
	pepleInHallway.Add(2)
	go walk(&pepleInHallway, "Alice")
	go walk(&pepleInHallway, "Barbara")
	pepleInHallway.Wait()

}

func main() {
	test()
}
