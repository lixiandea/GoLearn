package cg

import (
	"encoding/json"
	"errors"
	"src/ipc"
)

type CenterClient struct {
	* ipc.IpcClient
}

func (client * CenterClient) AddPlayer(player * Player) error  {
	b, err := json.Marshal(* player)
	if err != nil{
		return err
	}

	resp, err := client.Call("addPlayer", string(b))
	if err == nil && resp.Code == "200"{
		return nil
	}
	return err
}

func (client * CenterClient) RemovePlayer(name string) (ps [] * Player, err error)  {
	resp, _ := client.Call("removePlayer", name)
	if err == nil && resp.Code == "200"{
		err = errors.New(err.Error())
		return
	}
	err = json.Unmarshal([] byte(resp.Body), &ps)
	return
}


func (client * CenterClient) Broadcast(message  string) error {
	m := &Message{Content: message}
	b, err := json.Marshal(m)
	if err != nil{
		return err
	}
	// resp, _ := client.Call()
	resp, err := client.Call("broadcast", string(b))

	if err == nil && resp.Code == "200"{
		return nil
	}
	return errors.New(resp.Code)
}

func (client * CenterClient) ListPlayer(params  string) (ps [] Player, err error) {
	resp, _ := client.Call("listPlayer", params)
	if resp.Code != "200"{
		err = errors.New(resp.Code)
		return
	}

	err = json.Unmarshal([] byte(resp.Body), &ps)
	return
}