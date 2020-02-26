package main

// CommandInterface is ...
type CommandInterface interface {
	exec()
}

func exec(command CommandInterface) {
	command.exec()
}
