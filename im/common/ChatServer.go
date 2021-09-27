// package common
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

type Server struct {
	addr     string
	entering chan client
	leaving  chan client
	message  chan string
}

func NewChatServer(addr string) *Server {
	return &Server{
		addr:     addr,
		entering: make(chan client),
		leaving:  make(chan client),
		message:  make(chan string),
	}
}

func (s *Server) handleConn(conn net.Conn) {
	ch := make(chan string)

	who := conn.RemoteAddr().String()

	ch <- "you are " + who
	s.message <- who + "has entered"
	s.entering <- ch

	input := bufio.NewScanner(conn)

	go clientWrite(conn, ch)

	if input.Scan() {
		s.message <- who + ":" + input.Text()
	}

	s.leaving <- ch
	s.message <- who + " has leave"

}

func clientWrite(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func (s *Server) broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-s.message:
			for client := range clients {
				client <- msg
			}
		case client := <-s.entering:
			clients[client] = true
		case client := <-s.leaving:
			delete(clients, client)
			close(client)
		}
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.addr)

	if err != nil {
		log.Fatal("Server Start fail due to err:", err)
		panic(err)
	}
	go s.broadcast()

	log.Println("Server Start success and listen to ", s.addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer conn.Close()

		go s.handleConn(conn)
	}
}

func main() {
	server := NewChatServer(":10086")
	server.Start()
}
