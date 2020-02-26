package main

import "fmt"

// TurnOnCommand is ...
type TurnOnCommand struct{}

func (command TurnOnCommand) exec() {
	fmt.Println("The Light turns on")
}
