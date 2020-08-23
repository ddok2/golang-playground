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
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}
	return b
}
