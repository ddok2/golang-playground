// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - adder_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdder() {
	sum := Add(1, 5)
	fmt.Println(sum)
}
