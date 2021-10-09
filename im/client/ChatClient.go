// package common
package main

import (
	"bufio"
	"fmt"
	"im/common"
	"log"
	"net"
	"os"
	"sync"
)

var exitCh = make(chan string)

type Client struct {
	name string
	addr string
}

func setNickName() string {
	var nickname string
	fmt.Print("Enter your nickname : ")
	fmt.Scanln(&nickname)
	return nickname
}

func NewClient(addr string) *Client {
	return &Client{
		name:   setNickName(),
		addr: addr,
	}
}

func (c *Client) Start() {
	tcpaddr, err := net.ResolveTCPAddr("tcp", c.addr)
	if err != nil {
		log.Fatal("ResolveTCPAddr ", c.addr, " err: ", err)
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		log.Fatal("DialTCP ", c.addr, " err: ", err)
		panic(err)
	}

	defer conn.Close()
	// go c.handleSend(conn)

	go c.handleRecv(conn)
	var msg string
	for {
		msg = ""
		fmt.Scan(&msg)
		conn.Write([]byte(msg))
		if msg == "EXIT" {
			conn.Close()
		}
	}
}

func (c *Client) handleSend(conn net.Conn) {
	input := bufio.NewReader(os.Stdin)

	for{
		var msg common.Message
		var err error

		msg.User = c.name
		msg.Msg, err = input.ReadString('\n')
		if err != nil {
			fmt.Println("Stdin error ", err)
		}

		if msg.Msg[0] == '-' && msg.Msg[1] == '1' {
			fmt.Println("Exiting chatsession.")
			err := common.SendMsg(&common.Message{Msg: "leave Chat", User: c.name}, conn)
			if err != nil {
				log.Print("error send")
			}
			break
		}
		// Sends msg to chatserver.
		common.SendMsg(&msg, conn)
	}

}

func (c *Client) handleRecv(conn net.Conn) {
	for {
		msg, err := common.ReadMsg(conn)
		if err != nil {
			log.Println("Read fail due to err: ", err)
			continue
		}
		fmt.Println(msg.ToString())
	}
}

func main() {

	c := NewClient(":10086")
	go c.Start()
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
