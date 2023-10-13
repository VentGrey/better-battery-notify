/*
Package notify contains all the neccesary functions to trigger notifications,
as well as data structures and constants.
*/
package notify

import (
	"log"

	"github.com/godbus/dbus/v5"
)

// NotificationStatus represents the current state of battery notifications.
// It determines whether a particular notification has been shown to the user.
type NotificationStatus int

const (
	// NotNotified indicates that no notifications have been shown yet.
	NotNotified NotificationStatus = iota

	// LowBatteryNotified indicates that a low battery notification has been shown.
	LowBatteryNotified

	// VeryLowBatteryNotified indicates that a critical battery notification has been shown.
	VeryLowBatteryNotified
)

// SendNotification connects to the current session bus and tries to trigger
// a new notification via the freedesktop notifications interface.
func SendNotification(title, message, icon string) {
	conn, err := dbus.SessionBus()

	if err != nil {
		log.Println("Failed to connect to session bus: ", err)
		return
	}

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "", uint32(0),
		icon, title, message, []string{}, map[string]dbus.Variant{}, int32(5000))
	if call.Err != nil {
		log.Println("Failed to send notification:", call.Err)
	}
}
