// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - adder_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package integers

import (
	"fmt"
	"testing"
)

// func TestAdder(t *testing.T) {
// 	sum := Add(2, 2)
// 	expected := 4
//
// 	if sum != expected {
// 		t.Errorf("expected '%d' but got '%d'", expected, sum)
// 	}
// }
//
// func ExampleAdder() {
// 	sum := Add(1, 5)
// 	fmt.Println(sum)
// 	// Output: 6
// }

func TestAdd(t *testing.T) {

	f := make([]interface{}, 0)

	f = append(f, add("sung"))
	f = append(f, add("yub"))
	f = append(f, add("na"))

	lang := make([]string, 0)
	for _, v := range f {
		result := v.(func() string)()
		lang = append(lang, result)
	}

	fmt.Println(lang)

}

func add(args string) string {

	return args
}
