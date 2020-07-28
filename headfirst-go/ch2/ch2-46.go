// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - ch2-46.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	if input >= 60 {
		status := "passing"
	} else {
		status := "failing"
	}
}
