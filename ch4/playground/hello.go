// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - hello.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package hello

import (
	"fmt"

	"rsc.io/quote"
)

func Hello() string {
	return quote.Hello()
}

func Test1() {
	months := [...]string{
		1:  "Jan",
		2:  "Feb",
		3:  "Mar",
		4:  "Apr",
		5:  "May",
		6:  "Jun",
		7:  "Jul",
		8:  "Aug",
		9:  "Sep",
		10: "Oct",
		11: "Nov",
		12: "Dec",
	}

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println("Q2:", Q2)
	fmt.Println("summer:", summer)
	fmt.Println("months:", months)

	for _, s := range summer {
		for _, q := range Q2 {

			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
}
