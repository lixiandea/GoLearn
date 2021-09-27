package main

import (
	"fmt"
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "10086"
	CONN_TYPE = "tcp"
)

var connects []net.Conn

func main() {
	msgChan := make(chan string, 1024)
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer l.Close()

	fmt.Println("listening on: ", CONN_HOST+":"+CONN_PORT)

	go boradCastMessage(msgChan)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			panic(err)
		}
		defer conn.Close()
		connects = append(connects, conn)

		go handleRequest(conn, msgChan)
	}
}

func handleRequest(conn net.Conn, msgChan chan string) {
	buf := make([]byte, 1024)

	reqLen, err := conn.Read(buf)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	msgChan <- conn.RemoteAddr().String() + ":" + string(buf[:reqLen])
	fmt.Println(reqLen)
	// conn.Write([]byte("Message received."))
	// conn.Close()
}

func boradCastMessage(ch chan string) {
	var msg string
	for {
		select {
		case msg = <-ch:
			fmt.Println(msg)
			for _, conn := range connects {
				conn.Write([]byte(msg))
			}
		}
	}

}
