// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - main.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}

}

/*

### 테스트 수행 ###

$ go run main.go a bc def
a bc def

$ go run main.go -s / a bc def
a/bc/def

$ go run main.go -help
Usage of /var/folders/8p/bsxwcwr17kzgkwv1m9qr4crm0000gn/T/go-build442566164/b001/exe/main:
  -n    omit trailing newline
  -s string
        separator (default " ")
exit status 2
*/
