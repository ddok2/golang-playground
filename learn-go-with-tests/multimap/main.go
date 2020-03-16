// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - main.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"encoding/json"
	"fmt"

	"github.com/jwangsadinata/go-multimap/slicemultimap"
)

func main() {
	m := slicemultimap.New()

	// m.Put("elmo", 1000)
	// m.Put("elmo", 1010)
	// m.Put("elmo", 1020)
	// m.Put("elmo", 1030)

	balances, found := m.Get("elmo")

	if found {
		for i, balance := range balances {
			fmt.Println(i, ", ", balance)
		}

		fmt.Println(balances[len(balances)-1])
	}

	balances, found = m.Get("elmo2")

	if found {
		fmt.Println("found")
	}
	fmt.Println(balances, found)

	results := make(map[string]error)
	results["a"] = nil
	results["b"] = nil
	results["c"] = nil
	results["d"] = nil
	results["e"] = nil
	results["f"] = nil

	type Result struct {
		UUID string `json:"uuid"`
		Code string `json:"code"`
	}

	type ResultList struct {
		Results []*Result `json:"results"`
	}

	resultList := &ResultList{}

	for k, v := range results {
		r := &Result{}
		r.UUID = k
		if v == nil {
			r.Code = "OK"
		} else {
			r.Code = "ERROR"
		}
		resultList.Results = append(resultList.Results, r)
	}

	_json, _ := json.Marshal(resultList)
	fmt.Println(string(_json))
	fmt.Println(resultList)
}
