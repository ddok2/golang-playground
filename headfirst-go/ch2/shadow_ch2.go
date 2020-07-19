// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - shadow_ch2.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "fmt"

func main() {
	var count int = 12
	var suffix string = "minutes of bonus footage"
	var format string = "DVD"
	var languages = append([]string{}, "Korean")

	fmt.Println(count, suffix, "on", format, languages)
}
