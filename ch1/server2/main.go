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
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(writer, "URL.PATH = %q\n", request.URL.Path)
}

func counter(w http.ResponseWriter, _ *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
