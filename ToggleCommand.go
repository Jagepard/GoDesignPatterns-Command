package main

// ToggleCommand is ...
type ToggleCommand struct{}

var i int = 1

func (command ToggleCommand) exec() {
	if i == 1 {
		i--
		subCommand := TurnOnCommand{}
		subCommand.exec()
	} else {
		i++
		subCommand := TurnOffCommand{}
		subCommand.exec()
	}
}
