// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - racer_test.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package selelct

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("return an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

	// slowServer := makeDelayedServer(20 * time.Millisecond)
	// fastServer := makeDelayedServer(0 * time.Millisecond)
	//
	// slowURL := slowServer.URL
	// fastURL := fastServer.URL
	//
	// want := fastURL
	// got := Racer(slowURL, fastURL)
	//
	// if got != want {
	// 	t.Errorf("got %q, want %q", got, want)
	// }
	//
	// slowServer.Close()
	// fastServer.Close()
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		},
	))
}
