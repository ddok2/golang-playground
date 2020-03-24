// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - queue.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"container/list"
	"fmt"
	"sync"
)

type MessageQueue struct {
	rwMutex *sync.RWMutex
	queue   *list.List
}

var messageQueue *MessageQueue

func NewMessageQueue() *MessageQueue {
	if messageQueue == nil {
		messageQueue = new(MessageQueue)
	}

	return messageQueue
}

func (mq *MessageQueue) Initialize() {
	mq.queue = list.New()
	mq.rwMutex = new(sync.RWMutex)
}

func (mq *MessageQueue) SendMessage(msg string) {
	mq.rwMutex.Lock()
	defer mq.rwMutex.Unlock()
	mq.queue.PushBack(msg)
}

func (mq *MessageQueue) GetMessage() (string, error) {
	mq.rwMutex.Lock()
	defer mq.rwMutex.Unlock()
	if mq.queue.Len() > 0 {
		e := mq.queue.Front()
		msg := e.Value.(string)
		mq.queue.Remove(e)
		return msg, nil
	} else {
		return "", fmt.Errorf("empty queue")
	}
}

func (mq *MessageQueue) Finalize() {
	mq.rwMutex.Lock()
	defer mq.rwMutex.Unlock()

}
