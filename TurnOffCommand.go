/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

import "fmt"

// TurnOffCommand is ...
type TurnOffCommand struct{}

func (command TurnOffCommand) exec() {
	fmt.Println("The Light turns off")
}
