/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

// DeviceInterface is ...
type DeviceInterface interface {
	setCommand(commandName string, command CommandInterface)
	execute(commandName string)
}

func execute(device DeviceInterface, commandName string) {
	device.execute(commandName)
}

func setCommand(device DeviceInterface, commandName string, command CommandInterface) {
	device.setCommand(commandName, command)
}
