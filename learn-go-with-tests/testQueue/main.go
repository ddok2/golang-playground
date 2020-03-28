// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - main.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

import (
	"strings"
	"sync"
	"time"
)

var msgCount = 0

func main() {
	// startTime := time.Now()

	batchMessageCount := 0
	txBuffer := NewTxBuffer()

	for {
		if msg, err := messageQueue.GetMessage(); err != nil {
			time.Sleep(1 * time.Millisecond)
		} else {
			msgCount++
			// logger.Info("Received messages", msg)
			// logger.Info("ConsumeClaim goroutine id: ", getGID())
			aggregateTransaction(&batchMessageCount, msg, txBuffer)

		}

		// if c.checkSendTransaction(&batchMessageCount) {
		// 	c.sendTransaction(&batchMessageCount, txBuffer)
		// 	c.startTime = time.Now()
		// }
	}

}

// type TxBuffer struct {
//
// }

func aggregateTransaction(batchMessageCount *int, tx string, txBuffer *TxBuffer) {
	txArr := strings.Split(tx, ",")
	isDuplicate := false

	//if txArr[1] == RegisterMember {
	//	isDuplicate = !c.txKeyRegistry.Set(txArr[RegisterWalletOffset], tx)
	//} else if txArr[1] == TransferCoin {
	//	isDuplicate = !c.txKeyRegistry.Set(txArr[SenderWalletOffset], tx)
	//	//log.info("isDuplicate1: " + isDuplicate + ", txArr[2]: " + txArr[2])
	//	if !isDuplicate {
	//		isDuplicate = !c.txKeyRegistry.Set(txArr[ReceiverWalletOffset], tx)
	//		//log.info("isDuplicate2: " + isDuplicate + ", txArr[2]: " + txArr[2])
	//	}
	//}

	if isDuplicate {
		//if err := c.redisClient.Set(txArr[0], "400"); err != nil {
		//	logger.Errorf(err.Error())
		//	return
		//}

		//txProducer.runProducer(tx)
	} else {
		*batchMessageCount++
		//logger.Info("batchMessageCount : ", c.batchMessageCount)
		//logger.Info("txArr : ", txArr)

		if *batchMessageCount == 1 {
			txBuffer.Append(1, txArr...)
		} else {
			txBuffer.Append(0, txArr...)
		}

		//logger.Info("txArray : ", c.txArray)
	}
}

type TxBuffer struct {
	txArray []string
	rwMutex *sync.RWMutex
}

//var txBuffer *TxBuffer

func NewTxBuffer() *TxBuffer {
	//if txBuffer == nil {
	txBuffer := new(TxBuffer)
	txBuffer.rwMutex = new(sync.RWMutex)
	txBuffer.txArray = make([]string, 2, 1024)
	//}

	return txBuffer
}

func (tb *TxBuffer) Initialize(txBatchCount, maxBytesPerTx int) {
	tb.rwMutex.Lock()
	defer tb.rwMutex.Unlock()
	tb.txArray = nil
	tb.txArray = make([]string, 2, txBatchCount*maxBytesPerTx)
}

func (tb *TxBuffer) SetFirst(e string) {
	tb.rwMutex.Lock()
	defer tb.rwMutex.Unlock()
	tb.txArray[0] = e
}

func (tb *TxBuffer) Append(index int, txArr ...string) {
	tb.rwMutex.Lock()
	defer tb.rwMutex.Unlock()
	tb.txArray = append(tb.txArray[index:], txArr...)
}

func (tb *TxBuffer) Length() int {
	tb.rwMutex.Lock()
	defer tb.rwMutex.Unlock()

	return len(tb.txArray)
}

func (tb *TxBuffer) ToCopy() []string {
	tb.rwMutex.Lock()
	defer tb.rwMutex.Unlock()

	tempTxArray := make([]string, len(tb.txArray))
	copy(tempTxArray, tb.txArray)

	return tempTxArray
}
