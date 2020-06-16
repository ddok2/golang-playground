// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - gorutine_pool.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg, wg2 sync.WaitGroup

type Job struct {
	Id     int
	Work   string
	Result string
}

func produce(jobs chan<- *Job) {
	// Generate jobs:
	// id := 0
	// for c := 'a'; c <= 'z'; c++ {
	// 	id++
	// 	jobs <- &Job{Id: id, Work: fmt.Sprintf("%c", c)}
	// }

	jobs <- &Job{Id: 0, Work: "a"}
	jobs <- &Job{Id: 1, Work: "b"}
	jobs <- &Job{Id: 2, Work: "c"}
	jobs <- &Job{Id: 3, Work: "d"}
	jobs <- &Job{Id: 4, Work: "e"}
	jobs <- &Job{Id: 0, Work: "f"}
	jobs <- &Job{Id: 3, Work: "g"}
	jobs <- &Job{Id: 3, Work: "h"}
	jobs <- &Job{Id: 4, Work: "i"}

	close(jobs)
}

func consume(id int, jobs <-chan *Job, results chan<- *Job) {
	defer wg.Done()
	for job := range jobs {
		sleepMs := rand.Intn(1000)
		fmt.Printf("worker #%d received: '%s', sleep %dms\n", id, job.Work, sleepMs)
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
		job.Result = job.Work + fmt.Sprintf("-%dms", sleepMs)
		results <- job
	}
}

func analyze(results <-chan *Job) {
	defer wg2.Done()
	for job := range results {
		fmt.Printf("result: %s - %d \n", job.Result, job.Id)
	}
}

func main() {
	jobs := make(chan *Job, 100)    // Buffered channel
	results := make(chan *Job, 100) // Buffered channel

	// Start consumers:
	for i := 0; i < 5; i++ { // 5 consumers
		wg.Add(1)
		go consume(i, jobs, results)
	}
	// Start producing
	go produce(jobs)

	// Start analyzing:
	wg2.Add(1)
	go analyze(results)

	wg.Wait() // Wait all consumers to finish processing jobs

	// All jobs are processed, no more values will be sent on results:
	close(results)

	wg2.Wait() // Wait analyzer to analyze all results
}
