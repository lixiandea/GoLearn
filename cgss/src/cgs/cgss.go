package main

import (
	"bufio"
	"fmt"
	"os"
	"src/cg"
	"src/ipc"
	"strconv"
	"strings"
)

var centerClient * cg.CenterClient

func startCenterService() error {
	server := ipc.NewIpcServer(& cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &cg.CenterClient{IpcClient: client}
	return nil
}

func Help(args [] string) int {
	fmt.Println(`
	Commands:
		login <username> <level> <exp>
		logout <username> 
		send <message>
		listPlayer
		quit(q)
		help(h)
	`)
	return 0
}

func Quit(args [] string) int {
	return 1
}

func Logout(args [] string) int {
	if len(args) != 2{
		fmt.Println("USAGE: logout <username>")
		return 0
	}
	centerClient.RemovePlayer(args[1])
	return 0
}

func Login(args [] string) int {
	if len(args) != 4{
		fmt.Println("USAGE: login <username> <level> <exp>")
		return 0
	}
	level, err := strconv.Atoi(args[2])
	if err != nil{
		fmt.Println("Invalid parameter: <level> should be an integer")
		return 0
	}
	exp, err := strconv.Atoi(args[2])
	if err != nil{
		fmt.Println("Invalid parameter: <exp> should be an integer")
		return 0
	}

	player := cg.NewPlayer()
	player.Name = args[1]
	player.Level =  level
	player.Exp = exp
	centerClient.AddPlayer(player)
	return 0
}


func ListPlayer(args [] string) int {
	ps, err := centerClient.ListPlayer(" ")
	if err != nil{
		fmt.Println("Failed. ", err)
		return 0
	}else {
		for i,v := range ps{
			fmt.Println(i+1, ":", v)
		}
		return 0
	}
}

func Send(args[] string) int {
	message := strings.Join(args[1:], " ")
	err := centerClient.Broadcast(message)
	if err!=nil{
		fmt.Println("Failed :", err)
	}
	return 0
}

func handleCommandHandlers(args [] string) map[string] func([] string) int{
	return map[string]func([]string) int{
		"help": Help,
		"quit":Quit,
		"q": Quit,
		"h":Help,
		"login": Login,
		"logout": Logout,
		"listPlayer": ListPlayer,
		"l": ListPlayer,
		"send": Send,
	}

}

func main()  {
	fmt.Println("Casual Game server Solution.")
	startCenterService()

	Help(nil)

	r := bufio.NewReader(os.Stdin)

	handlers := handleCommandHandlers(nil)
	for{
		fmt.Print("command->")
		b, _, _ := r.ReadLine()
		line := string(b)

		args := strings.Split(line, " ")
		if handler, ok := handlers[args[0]]; ok{
			ret := handler(args)
			if ret != 0{
				break
			}
		}else {
			fmt.Println("Unknown command: ", args[0])
		}
	}
}

