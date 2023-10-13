/*
Package battery contains all the functions needed to process battery outputs
and monitoring.
*/
package battery

import (
	"battery-notify/notify"
	"battery-notify/system"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fsnotify/fsnotify"
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

// MonitorBattery contains the main battery monitoring logic, using fsnotify
// to get new battery status/values.
func MonitorBattery(notified *notify.NotificationStatus, battery_path string, dead, ideal_charge_min, ideal_charge_max uint) {
	
	// Initialize a new watcher
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()


	// Add the battery (from path) to the watcher
	err = watcher.Add(battery_path)

	if err != nil {
		log.Fatal(err)
	}

	// Main process loop (spawn in a goroutine)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					batteryStatus := ReadBatteryStatus(battery_path)
					batteryLevel := ReadBatteryLevel(battery_path)

					switch batteryStatus {
					case "Discharging":
						if batteryLevel <= dead && *notified != notify.VeryLowBatteryNotified {
							notify.SendNotification("Battery Critical", "Suspending in a few moments", "battery")
							*notified = notify.VeryLowBatteryNotified
							system.SuspendSystem()
						} else if batteryLevel <= ideal_charge_min && *notified != notify.LowBatteryNotified {
							notify.SendNotification("Battery Low", "Consider plugging in for battery longevity", "battery")
							*notified = notify.LowBatteryNotified
						}
					case "Charging":
						if batteryLevel >= ideal_charge_max && *notified != notify.NotNotified {
							notify.SendNotification("Battery Almost Full", "Consider unplugging for battery longevity", "battery")
							*notified = notify.NotNotified
						}
					case "Full":
						if batteryLevel == 100 && *notified != notify.NotNotified {
							notify.SendNotification("Battery Full", "Consider unplugging", "battery")
							*notified = notify.NotNotified
						}
					default:
						log.Println("Unknown battery status:", batteryStatus)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	select{}
}
