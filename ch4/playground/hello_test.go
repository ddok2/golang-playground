// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - hello_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."

	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

// func TestTest1(t *testing.T) {
// 	Test1()
// }

/**
 go test
go: finding rsc.io/quote v1.5.2
go: downloading rsc.io/quote v1.5.2
go: extracting rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: extracting rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
PASS
ok      playground/hello/playground     0.300s
*/
