// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - racer.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package selelct

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func measureReponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
