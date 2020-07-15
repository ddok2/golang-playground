// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - pass_fail.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}
