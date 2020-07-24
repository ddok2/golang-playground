// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - dictionary.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
