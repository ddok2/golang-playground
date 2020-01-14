package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen(`tcp`, `localhost:3000`)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format(`15:04:05`+"\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}

/**

** 실행 방법 **

$ export GOPATH=$(pwd)
$ go get github.com/ddok2/golang-playground/ch8/clock1
$ ./bin/clock1 &
$ nc localhost 3000
19:47:53
19:47:54
19:47:55
19:47:56
19:47:57
19:47:58
^C

**/
