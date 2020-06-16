// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - gorutine_pool3.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import "fmt"

type imageMessage struct {
	path string
	img  string
}

var paths []string
var images []string

func init() {
	for i := 0; i < 26; i++ {
		paths = append(paths, fmt.Sprintf("/img/%c", 'a'+i))
		images = append(images, fmt.Sprintf("%c.png", 'a'+i))
	}
}

func processImage(path, img string) {
	fmt.Printf("Processing %s/%s\n", path, img)
}

func worker(jobs <-chan imageMessage, results chan<- imageMessage) {
	for j := range jobs {
		processImage(j.path, j.img)
		results <- j
	}
}

func main() {
	jobs := make(chan imageMessage, 1)
	results := make(chan imageMessage, 1)

	for w := 0; w < 2; w++ {
		go worker(jobs, results)
	}

	go func() {
		for j := 0; j < len(images); j++ {
			jobs <- imageMessage{path: paths[j], img: images[j]}
		}
		close(jobs)
	}()

	for r := 0; r < len(images); r++ {
		<-results
	}
}
