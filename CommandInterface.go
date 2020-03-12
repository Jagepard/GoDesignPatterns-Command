/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

// CommandInterface is ...
type CommandInterface interface {
	exec()
}

func exec(command CommandInterface) {
	command.exec()
}
