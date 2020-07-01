// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - hash.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"fmt"
	"hash/fnv"
)

func hash(walletId string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(walletId))

	return h.Sum32()
}

func main() {
	walletId := "elmo"

	queueId := hash(walletId)

	fmt.Println(queueId % 10)
}
