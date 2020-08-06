// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - countdown.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}
