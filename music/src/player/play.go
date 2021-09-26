package player

import "fmt"

type Player interface {
	Play(source string)
}


func Play(source, musicType string){
	var p Player

	switch musicType {
	case "MP3":
		p = &MP3Player{}
	default:
		fmt.Println("not support format: ", musicType)
	}
	p.Play(source)
}
