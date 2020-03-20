// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - main.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jwangsadinata/go-multimap/slicemultimap"
)

type Result struct {
	UUID string `json:"uuid"`
	Code string `json:"code"`
}

type ResultList struct {
	Results []*Result `json:"results"`
}

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

	// val, found := results["z"]
	// if found {
	// 	fmt.Println("found key of z", val)
	// } else {
	// 	fmt.Println(results["z"])
	// }

	resultList := &ResultList{}

	resultList.Results = append(resultList.Results, &Result{
		UUID: "a",
		Code: "1",
	})
	resultList.Results = append(resultList.Results, &Result{
		UUID: "b",
		Code: "2",
	})

	resultList.Results = append(resultList.Results, test1().Results...)

	// for k, v := range results {
	// 	r := &Result{}
	// 	r.UUID = k
	// 	if v == nil {
	// 		r.Code = "OK"
	// 	} else {
	// 		r.Code = "ERROR"
	// 	}
	// 	resultList.Results = append(resultList.Results, r)
	// }
	//
	// _json, _ := json.Marshal(resultList)
	// fmt.Println(string(_json))
	// fmt.Println(resultList)

	t := test2([]string{"a", "b", "c"})

	for _, v := range t.Results {
		if strings.Contains(v.Code, "elmo") {
			fmt.Println(v)
		}
	}

	_json, _ := json.Marshal(t)
	fmt.Println(string(_json))

	fmt.Println(string(http.StatusInternalServerError))
	fmt.Println(http.StatusText(http.StatusInternalServerError))
}

func test1() *ResultList {
	resultList := &ResultList{}
	resultList.Results = append(resultList.Results, &Result{
		UUID: "c",
		Code: "1",
	})
	resultList.Results = append(resultList.Results, &Result{
		UUID: "d",
		Code: "2",
	})

	return resultList
}

func test2(str []string) *ResultList {
	resultList := &ResultList{}

	// str[0] = a
	// str[1] = b
	// str[2] = c

	for i, v := range str {

		switch v {
		case "a":
			fmt.Println("a")
			r := &Result{
				UUID: "a",
				Code: "ok",
			}

			resultList.Results = append(resultList.Results, r)

			if i == 0 {
				r.Code = "err"
				continue
			}

		case "b":
			fmt.Println("b")
			r := &Result{
				UUID: "b",
				Code: "elmo",
			}
			resultList.Results = append(resultList.Results, r)
			if i == 0 {
				r.Code = "err"
				continue
			}

		case "c":
			fmt.Println("c")
			r := &Result{
				UUID: "c",
				Code: "something elmo whatever",
			}
			resultList.Results = append(resultList.Results, r)
			if i == 0 {
				r.Code = "err"
				continue
			}

		}

	}

	return resultList
}
