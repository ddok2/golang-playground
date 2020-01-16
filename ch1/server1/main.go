/*
 * Copyright 2020. SUNGYUB.COM. All Rights Reserved.
 *
 * - main.go
 * - author: Sungyub NA <mailto: darkerkorean@gmail.com>
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // 각 요청은 핸들러를 호출
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

// 요청된 URL r의 Path 구성요소를 리턴
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
}
