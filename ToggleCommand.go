package main

// ToggleCommand is ...
type ToggleCommand struct{}

func (command ToggleCommand) exec() {
	toggle := 0

	if toggle == 1 {
		toggle--
		subCommand := TurnOnCommand{}
		subCommand.exec()
	} else {
		toggle++
		subCommand := TurnOffCommand{}
		subCommand.exec()
	}
}
