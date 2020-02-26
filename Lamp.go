package main

// Lamp is ...
type Lamp struct {
	commands map[string]CommandInterface
}

func (decice Lamp) setCommand(commandName string, command CommandInterface) {
	decice.commands[commandName] = command
}

func (decice Lamp) execute(commandName string) {
	decice.commands[commandName].exec()
}
