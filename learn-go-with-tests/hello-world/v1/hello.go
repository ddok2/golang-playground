// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - hello.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

const englishHelloPrefix string = "Hello, "
const spanishHelloPrefix string = "Hola, "

func Hello(name string, lang string) string {

	if name == "" {
		name = "World"
	}

	return setLang(lang) + name
}

func setLang(lang string) string {
	switch lang {
	case "Spanish":
		return spanishHelloPrefix
	default:
		return englishHelloPrefix
	}
}
