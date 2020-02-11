// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - main.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"bufio"
	"fmt"
	"os"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(comma(input.Text()))
	}
}

/**
sadfsdfsdf
s,adf,sdf,sdf
sdfsadf
s,dfs,adf
as
as
ass
ass
sssss
ss,sss
ssssdfasdfsdaf
ss,ssd,fas,dfs,daf
sadfsadfsadfdsafdsafsadf
sad,fsa,dfs,adf,dsa,fds,afs,adf
*/
