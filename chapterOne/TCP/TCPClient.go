package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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
	go readConn(conn)
	go readConn(conn)
	select {}
}

func readConn(conn net.Conn) {
	for {
		reply := make([]byte, 1024)

		n, err := conn.Read(reply)

		if err != nil {
			log.Fatal("read %s, err: %s", conn.RemoteAddr(), err)
			panic(err)
		}
		if n != 0 {
			fmt.Println("receive message: ", string(reply[:n]))
		}
	}
}

func writeConn(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(">>")
		text, _ := reader.ReadString('\n')
		data := []byte(text)
		conn.Write(data)
	}
}
