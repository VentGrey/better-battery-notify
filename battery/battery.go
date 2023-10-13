/*
Package battery contains all the functions needed to process battery outputs
and monitoring.
*/
package battery

import (
	"battery-notify/battery"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadBatteryStatus gets the current battery status number from the provided
// battery status file at main.
func ReadBatteryStatus(battery_path string) string {
	status, err := os.ReadFile(battery_path + "/status")

	if err != nil {
		log.Println("Error reading battery status: ", err)
		return ""
	}
	return strings.TrimSpace(string(status))
}

// ReadBatteryLevel gets the current battery level from the provided
// battery path file at main.
func ReadBatteryLevel(battery_path string) uint {
	levelStr, err := os.ReadFile(battery_path + "/capacity")

	if err != nil {
		log.Println("Error reading battery level: ", err)
		return 0
	}

	level, err := strconv.Atoi(strings.TrimSpace(string(levelStr)))

	if err != nil {
		log.Println("Error converting battery level to integer: ", err)
		return 0
	}
	return uint(level)
}
