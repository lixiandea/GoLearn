package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	strEcho := "Hello"
	servAddr := "localhost:10086"

	tcpaddr, err := net.ResolveTCPAddr("tcp", servAddr)

	if err != nil {
		log.Fatal("ResolveTCPAddr %s, err: %s", servAddr, err)
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		log.Fatal("DialTCP %s, err: %s", tcpaddr, err)
		panic(err)
	}

	defer conn.Close()

	_, err = conn.Write([]byte(strEcho))

	if err != nil {
		log.Fatal("Write %s, err: %s", strEcho, err)
		panic(err)
	}

	fmt.Println("write to server = ", strEcho)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)

	if err != nil {
		log.Fatal("read %s, err: %s", tcpaddr, err)
		panic(err)
	}

	fmt.Println("receive message: ", string(reply))

	select {}
}
