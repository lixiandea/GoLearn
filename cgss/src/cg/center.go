package cg

import (
	"encoding/json"
	_ "encoding/json"
	"errors"
	"src/ipc"
	"sync"
)
type Message struct {
	From string `json:"from"`
	To string `json:"to"`
	Content string `json:"content"`

}

var _ ipc.Server

type CenterServer struct {
	servers map[string] ipc.Server
	players [] * Player
	rooms [] *Room
	mutex sync.Mutex
}

func NewCenterServer() * CenterServer  {
	servers := make(map[string] ipc.Server)
	players := make([] * Player, 0)
	return  & CenterServer{
		servers: servers,
		players: players,
	}
}

func (server * CenterServer) addPlayer(params string) error {
	player := NewPlayer()
	err := json.Unmarshal([]byte(params), & player)
	if err != nil{
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()

	server.players = append(server.players, player)
	return nil
}

func (server * CenterServer)removePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for i,v := range server.players{
		if v.Name == params{
			if len(server.players) == 1{
				server.players = make([] * Player, 0)
			}else if i == len(server.players) - 1{
				server.players = server.players[:i]
			}else if i == 0 {
				server.players = server.players[1:]
			}else {
				server.players = append(server.players[:i], server.players[i+1 :]...)
			}
			return nil
		}
	}
	return errors.New("Player not found. ")
}

func (server * CenterServer) listPlayer(params string) (players string, err error) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	if len(server.players) > 0 {
		b, _ := json.Marshal(server.players)
		players = string(b)
	}else {
		err = errors.New("No player online! ")
	}
	return
}

func (server * CenterServer) broadcast(params string) error {
	var message Message
	err := json.Unmarshal([] byte(params), & message)
	if err != nil{
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	if len(server.players) > 0{
		for _, player := range server.players{
			player.mq <- &message
		}
	}else {
		err = errors.New("No player online. ")
	}
	return err
}

func (server * CenterServer)Handle(method, params string) * ipc.Response {
	switch method {
	case "addPlayer":
		err := server.addPlayer(params)
		if err !=nil{
			return &ipc.Response{Code: err.Error()}
		}
	case "removePlayer":
		err := server.removePlayer(params)
		if err != nil{
			return &ipc.Response{Code: err.Error()}
		}
	case "listPlayer":
		players, err := server.listPlayer(params)
		if err != nil{
			return &ipc.Response{Code: err.Error()}
		}
		return  & ipc.Response{Code: "200", Body: players}
	case "broadcast":
		err := server.broadcast(params)
		if err != nil{
			return & ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{Code: "200"}
	default:
		return &ipc.Response{Code: "404",Body: method + ": " + params}
	}
	return &ipc.Response{Code: "200"}
}

func (server * CenterServer)Name() string {
	return "CenterServer"

}

