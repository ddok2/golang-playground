// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - forcontinue_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package testFor

import (
	"testing"
)

func Test_frr(t *testing.T) {
	want := []string{
		"user2",
		"1",
		"2",
		"3",
	}
	got := frr()

	for i, v := range got {

		switch v {
		case want[i]:
			break
		default:
			t.Errorf("frr() = %v, want %v", v, want[i])
		}
	}
}
