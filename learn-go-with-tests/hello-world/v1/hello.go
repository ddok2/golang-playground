// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - hello.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

const englishHelloPrefix string = "Hello, "

func Hello(name string) string {

	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}
