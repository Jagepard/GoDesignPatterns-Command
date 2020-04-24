/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

func main() {
	device := Lamp{commands: make(map[string]CommandInterface)}

	device.setCommand("on", TurnOnCommand{})
	device.setCommand("off", TurnOffCommand{})
	device.setCommand("toggle", ToggleCommand{})

	device.execute("on")
	device.execute("off")
	device.execute("toggle")
	device.execute("toggle")
}
