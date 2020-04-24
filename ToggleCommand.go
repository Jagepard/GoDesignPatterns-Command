/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

// ToggleCommand is ...
type ToggleCommand struct{}

var toggle int = 0

func (command ToggleCommand) exec() {
	if toggle == 1 {
		subCommand := TurnOffCommand{}
		subCommand.exec()
		toggle--
	} else {
		subCommand := TurnOnCommand{}
		subCommand.exec()
		toggle++
	}
}
