// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - forcontinue.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package testFor

import "fmt"

func frr() {
	strs := []string{
		"elmo",
		"user1",
		"user2",
		"1",
		"2",
		"3",
	}

	newStrs := make([]string, 0)
	for i, str := range strs {
		if str == "elmo" {
			continue
		}
		if str == "user1" {
			continue
		}
		newStrs = append(newStrs, string(i)+str)
	}

	fmt.Println(newStrs)
}
