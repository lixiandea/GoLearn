// package common
package main

import (
	"fmt"
	"im/common"
	"log"
	"net"
)

type client chan<- string

type Server struct {
	addr    string
	clients []net.Conn
}

func NewChatServer(addr string) *Server {
	var clients [] net.Conn
	return &Server{
		addr:    addr,
		clients: clients,
	}
}

func (s *Server) deleteClosedConn(conn net.Conn) {
	// Finds the connection in clients and removes it.
	for i := 0; i <= len(s.clients); i++ {
		if conn == s.clients[i] {
			s.clients = s.clients[:i+copy(s.clients[i:], s.clients[i+1:])]
			break
		}
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		msg, err := common.ReadMsg(conn)
		//The connection got closed.
		if err != nil {
			fmt.Println("Lost connection to client")
			s.deleteClosedConn(conn)
			break
		}

		fmt.Println("Recived msg : ", msg.Msg, " from ", msg.User)
		// Sends msg to all clients connected to server.
		for _, conn := range s.clients {
			common.SendMsg(&msg, conn)
		}
	}

}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.addr)

	if err != nil {
		log.Fatal("Server Start fail due to err:", err)
		panic(err)
	}

	log.Println("Server Start success and listen to ", s.addr)

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}else {
			s.clients = append(s.clients, conn)
			go s.handleConn(conn)
		}

	}
}

func main() {
	server := NewChatServer(":10086")
	go server.Start()
	var input string
	fmt.Println("Exit server by pressing enter in console.")
	fmt.Scanln(&input)
	fmt.Println("Exiting server")
}
