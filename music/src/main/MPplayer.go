package main

import (
	"bufio"
	"fmt"
	"os"
	"src/musicLib"
	"src/player"
	"strconv"
	"strings"
)

var lib *musicLib.MusicManager
var id int = 1
var ctrl, signal chan int
var mp player.Player

func handleLibCommand(token []string) {
	switch token[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			m, _ := lib.Get(i)
			fmt.Println(i+1, m.Id, m.Name, m.Type, m.Artist, m.Source, m.Type)
		}
	case "add":
		if len(token) == 6 {
			lib.Add(&musicLib.MusicEntry{Id: strconv.Itoa(id), Name: token[2], Artist: token[3], Source: token[4], Type: token[5]})
			id++
		} else {
			fmt.Println("USAGE: lib add <name> <artist> <source> <type>")
		}
	case "remove":
		if len(token) == 3 {
			lib.RemoveByName(token[2])
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command: ", token[1])
	}
}

func handlePlayerCommand(token []string) {
	if len(token) != 2 {
		fmt.Println("USAGE ： play name")
		return
	} else {
		e, _ := lib.Find(&token[1])
		if e == nil {
			fmt.Println("The Music ", token[1], "not found")
			return
		}
		player.Play(e.Source, e.Type)
	}
}

func main() {
	fmt.Println(`
		Enter following commands to control the player:
		lib list -- view the existing music lib
		lib add <name> <artist> <source> <type>  -- add a music to the music lib
		lib remove <name> -- remove the specified music from the lib
		play <name> -- play the specified name
	`)

	lib = musicLib.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter Command ->")
		rawline, _, _ := r.ReadLine()
		line := string(rawline)
		if line == "q" || line == "e" {
			break
		}
		args := strings.Split(line, " ")

		switch args[0] {
		case "lib":
			handleLibCommand(args)
		case "play":
			handlePlayerCommand(args)
		default:
			fmt.Println("wrong command :", args[0])
		}
	}

}
