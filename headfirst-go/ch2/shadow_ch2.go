// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - shadow_ch2.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "fmt"

func main() {
	var count = 12
	var suffix = "minutes of bonus footage"
	var format = "DVD"
	var languages = append([]string{}, "Korean")

	fmt.Println(count, suffix, "on", format, languages)
}
