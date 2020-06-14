// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - gorutine_pool2.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg3 sync.WaitGroup

type Job2 struct {
	Id   int
	Work string
}

func produce2(jobs chan<- *Job2) {
	// Generate jobs:
	id := 0
	for c := 'a'; c <= 'z'; c++ {
		id++
		jobs <- &Job2{Id: id, Work: fmt.Sprintf("%c", c)}
	}
	close(jobs)
}

func consume2(id int, jobs <-chan *Job2) {
	defer wg3.Done()
	for job := range jobs {
		sleepMs := rand.Intn(1000)
		fmt.Printf("worker #%d received: '%s', sleep %dms\n", id, job.Work, sleepMs)
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
		fmt.Printf("result: %s\n", job.Work+fmt.Sprintf("-%dms", sleepMs))
	}
}

func main() {
	jobs := make(chan *Job2, 100) // Buffered channel

	// Start consumers:
	for i := 0; i < 5; i++ { // 5 consumers
		wg3.Add(1)
		go consume2(i, jobs)
	}
	// Start producing
	go produce2(jobs)

	wg3.Wait() // Wait all consumers to finish processing jobs
}
