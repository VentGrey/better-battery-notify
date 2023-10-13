/*
Package system contains all utilities to interact with the user's system and
session bus for various tasks.
*/
package system

import (
	"log"

	"github.com/godbus/dbus/v5"
)

// SuspendSystem uses the system bus to send a suspension signal to the system
// when the battery level reaches "dead" as defined in the main package.
func SuspendSystem() {
	conn, err := dbus.SystemBus()

	if err != nil {
		log.Println("Failed to connect to system bus:", err)
		return
	}
	obj := conn.Object("org.freedesktop.login1", "/org/freedesktop/login1")
	call := obj.Call("org.freedesktop.login1.Manager.Suspend", 0, true)
	if call.Err != nil {
		log.Println("Failed to suspend the system:", call.Err)
	}
}
