// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - adder_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package integers

import (
	"fmt"
	"reflect"
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

// func TestAdd(t *testing.T) {
//
// 	f := make([]interface{}, 0)
//
// 	f = append(f, add("sung"))
// 	f = append(f, add("yub"))
// 	f = append(f, add("na"))
//
// 	lang := make([]string, 0)
// 	for _, v := range f {
// 		result := v.(func() string)()
// 		lang = append(lang, result)
// 	}
//
// 	fmt.Println(lang)
//
// }
//
// func add(args string) string {
//
// 	return args
// }

func TestMulti(t *testing.T) {

	testArr := make([]interface{}, 0)

	testArr = append(testArr, &typeA{
		name:    "a",
		number1: 1,
		number2: 2,
	})

	testArr = append(testArr, &typeB{
		name:    "B1",
		number:  2,
		number2: 0,
	})

	testArr = append(testArr, &typeB{
		name:    "B2",
		number:  4,
		number2: 0,
	})

	for i, v := range testArr {
		t := reflect.TypeOf(v).String()
		switch t {
		case "*integers.typeA":
			fmt.Println("this is a")
			fmt.Println(v.(*typeA).Add())
			fmt.Println(i)

			break

		case "*integers.typeB":
			fmt.Println("this is b")
			fmt.Println(v)
			// fmt.Println(testArr[i].(typeB).Sub())
			break
		}
	}
}

type typeA struct {
	name    string
	number1 int
	number2 int
}

func (a *typeA) Add() int {
	return a.number1 + a.number2
}

type typeB struct {
	name    string
	number  int
	number2 int
}

func (b *typeB) Sub() int {
	return b.number2 - b.number
}
