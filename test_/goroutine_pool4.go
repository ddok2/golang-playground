// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - goroutine_pool4.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"fmt"
	"sync"
)

type Data struct {
	Flag       bool
	Id         string
	WalletAddr string
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(5)

	p := sync.Pool{
		New: func() interface{} {
			return []Data{
				{Flag: false, Id: "a", WalletAddr: "a"},
				{Flag: false, Id: "b", WalletAddr: "b"},
				{Flag: false, Id: "c", WalletAddr: "c"},
				{Flag: false, Id: "d", WalletAddr: "d"},
				{Flag: false, Id: "e", WalletAddr: "e"},
			}
		},
	}

	for n := 0; n < 5; n++ {
		goId := n
		go func() {
			slice := p.Get().([]Data)

			fmt.Println(slice)

			newSlice := make([]Data, 0)

			for i := 0; i < len(slice); i++ {
				if string('a'+goId) == slice[i].Id {
					newSlice = append(newSlice, slice[i])
					slice[i].Flag = true
				}
			}

			p.Put(slice)
			fmt.Println("goID", goId, " : ", slice)

			wg.Done()
		}()
	}

	wg.Wait()
}
