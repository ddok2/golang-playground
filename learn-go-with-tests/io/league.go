// Copyright 2021. SUNGYUB.COM. All Rights Reserved.
//
// - league.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}
