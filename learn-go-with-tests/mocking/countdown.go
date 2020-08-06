// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - countdown.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
