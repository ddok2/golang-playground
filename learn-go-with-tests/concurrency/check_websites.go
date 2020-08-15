// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - check_websites.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func() {
			results[url] = wc(url)
		}()
	}
	time.Sleep(2 * time.Second)

	return results
}
