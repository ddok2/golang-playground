// Copyright 2021. SUNGYUB.COM. All Rights Reserved.
//
// - sync.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package sync

type Counter struct {
	value int
}

func (c *Counter) Inc() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
