package common

import (
	"encoding/gob"
	"log"
	"net"
)

func SendMsg(msg *Message, conn net.Conn) error{
	encoder := gob.NewEncoder(conn)
	err := encoder.Encode(msg)
	if err != nil {
		log.Print("encoder fail, err:", err)
	}
	return err
}

func ReadMsg(conn net.Conn) (Message, error){
	var msg Message
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(&msg)
	if err != nil {
		log.Print("encoder fail, err:", err)
	}
	// fmt.Print(msg.ToString())
	return msg, err
}
