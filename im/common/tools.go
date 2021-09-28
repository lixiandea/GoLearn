package common

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

func sendMsg(msg *Message, conn net.Conn) {
	encoder := gob.NewEncoder(conn)
	err := encoder.Encode(msg)
	if err != nil {
		log.Print("encoder fail, err:", err)
	}
}

func readMsg(msg *Message, conn net.Conn) {
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(msg)
	if err != nil {
		log.Print("encoder fail, err:", err)
	}
	fmt.Print(msg.ToString())
}
