/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

func main() {
	device := Lamp{commands: make(map[string]CommandInterface)}

	setCommand(device, "on", TurnOnCommand{})
	setCommand(device, "off", TurnOffCommand{})
	setCommand(device, "toggle", ToggleCommand{})

	execute(device, "on")
	execute(device, "off")
	execute(device, "toggle")
	execute(device, "toggle")
}
