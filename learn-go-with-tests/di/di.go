// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - di.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
