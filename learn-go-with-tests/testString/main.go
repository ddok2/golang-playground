// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - main.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"fmt"
	"strings"
)

func main() {

	// elmoBalance := 0

	output := "Incorrect Balance, should:9999990000, not:10000000000"

	if strings.Contains(output, "Incorrect") {
		test1 := strings.Split(output, ", ")
		test2 := strings.Split(test1[1], ":")[1]
		fmt.Println(test2)
	}

}
