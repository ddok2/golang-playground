// Copyright 2021. SUNGYUB.COM. All Rights Reserved.
//
// - sync_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package sync

import "testing"

func assertCounter(t testing.TB, got Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3",
		func(t *testing.T) {
			counter := Counter{}
			counter.Inc()
			counter.Inc()
			counter.Inc()

			assertCounter(t, counter, 3)
		})
}
