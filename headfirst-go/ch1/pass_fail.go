// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - pass_fail.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
