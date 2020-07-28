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
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		fmt.Println("passing")
	} else {
		fmt.Println("failing")
	}
}
