// Copyright 2021. SUNGYUB.COM. All Rights Reserved.
//
// - file_system_store.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(f.database)
	return league
}
