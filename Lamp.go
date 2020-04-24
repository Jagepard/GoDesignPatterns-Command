/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

// Lamp is ...
type Lamp struct {
	commands map[string]CommandInterface
}

func (device Lamp) setCommand(commandName string, command CommandInterface) {
	device.commands[commandName] = command
}

func (device Lamp) execute(commandName string) {
	device.commands[commandName].exec()
}
