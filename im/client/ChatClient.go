// package common
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var exitCh = make(chan string)

type Client struct {
	ch   chan string
	addr string
}

func NewClient(addr string) *Client {
	return &Client{
		ch:   make(chan string, 1024),
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
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if len(line) != 0 {
			fmt.Fprintln(conn, line)
		}
		if line == "exit" {
			exitCh <- "exit"
		}
		for t := range c.ch {
			fmt.Println(t)
		}
	}
}

func (c *Client) handleRecv(conn net.Conn) {
	buff := make([]byte, 4096)
	len := 0
	for {
		n, err := conn.Read(buff)
		if err != nil {
			log.Println("Read fail due to err: ", err)
			continue
		}

		if n == 0 {
			continue
		}
		fmt.Println(buff[:len])
	}
}

func main() {
	c := NewClient(":10086")
	c.Start()
}
