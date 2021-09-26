package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	Addr := "localhost:10087"
	s, err := net.ResolveUDPAddr("udp4", Addr)

	if err != nil {
		log.Fatal("ResolveUDPAddr %s, err: %s", Addr, err)
		panic(err)
	}

	c, err := net.DialUDP("udp4", nil, s)

	if err != nil {
		log.Fatal("DialUDP %s, err: %s", Addr, err)
		panic(err)
	}

	fmt.Printf("connect to udp server %s", c.RemoteAddr().String())

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(">>")

		text, _ := reader.ReadString('\n')
		data := []byte(text)

		_, err = c.Write(data)
		if err != nil {
			log.Fatal("write failed, err : ", err)
			continue
		}
		if strings.TrimSpace(text) == "STOP" {
			fmt.Println("Existing UDP Client!")
			return
		}

		buff := make([]byte, 1024)

		n, _, err := c.ReadFromUDP(buff)

		if err != nil {
			log.Fatal("readFromUDP failed, err :", err)
			continue
		}
		fmt.Println("Reply: %s\n", string(buff[:n]))
	}
}
