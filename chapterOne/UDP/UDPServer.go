package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	PORT = ":10087"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	// udp4 means
	udpAdd, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		log.Fatal("ResolveUDPAddr %s, err: %s", udpAdd, err)
		panic(err)
	}

	conn, err := net.ListenUDP("udp4", udpAdd)
	if err != nil {
		log.Fatal("ListenUDP %s, err: %s", udpAdd, err)
		panic(err)
	}
	defer conn.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal("ReadFromUDP failed, err :", err)
			continue
		}

		fmt.Print("-> ", string(buffer[:n-1]))
		// 最后一个字符会有/r的问题，所以是n-1
		if strings.TrimSpace(string(buffer[:n-1])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		data := []byte(strconv.Itoa(random(1, 1001)))

		fmt.Printf("data: %s\n", string(data))

		_, err = conn.WriteToUDP(data, addr)
		if err != nil {
			log.Fatal("WriteToUDP failed, err :", err)
		}
	}
}
