package main

import "fmt"

// TurnOffCommand is ...
type TurnOffCommand struct{}

func (command TurnOffCommand) exec() {
	fmt.Println("The Light turns off")
}
