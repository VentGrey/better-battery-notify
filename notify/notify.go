/*
Package notify contains all the neccesary functions to trigger notifications,
as well as data structures and constants.
*/
package notify

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
