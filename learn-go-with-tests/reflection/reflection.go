// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - reflection.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
