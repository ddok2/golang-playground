// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - dictionary.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	return d[word], nil
}
