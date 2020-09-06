// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - reflection.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

func walk(x interface{}, fn func(input string)) {
	fn("whatever")
}
