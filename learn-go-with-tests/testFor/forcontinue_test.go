// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - forcontinue_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package testFor

import (
	"fmt"
	"testing"
)

// func Test_frr(t *testing.T) {
// 	want := []string{
// 		"user2",
// 		"1",
// 		"2",
// 		"3",
// 	}
// 	got := frr()
//
// 	for i, v := range got {
//
// 		switch v {
// 		case want[i]:
// 			break
// 		default:
// 			t.Errorf("frr() = %v, want %v", v, want[i])
// 		}
// 	}
// }

func TestSlice(t *testing.T) {
	type a struct {
		name  string
		value int
	}

	type b struct {
		name     string
		value    int
		whatever int
	}

	testList := make([]interface{}, 0)
	testList = append(testList, a{
		name:  "1",
		value: 1,
	})
	testList = append(testList, a{
		name:  "2",
		value: 2,
	})
	testList = append(testList, b{
		name:     "1",
		value:    1,
		whatever: 1,
	})
	testList = append(testList, b{
		name:     "2",
		value:    2,
		whatever: 2,
	})
	testList = append(testList, a{
		name:  "3",
		value: 3,
	})

	for i, v := range testList {
		switch t := v.(type) {
		case a:
			fmt.Println("this is a")
			fmt.Println(i, v)
		case b:
			fmt.Println("this is b")
			fmt.Println(i, v)

		default:
			fmt.Println("Unsupported type:", t)
		}

	}
}
