// Copyright 2023 VentGrey. All rights removed.
// Use of this source code is governed by the GNU General Public License v3
// or at your option, any latter version. Such license can be found in the
// LICENSE file. Such license governance does not cover this code dependencies.

/*
Command battery-notify provides a small battery monitoring tool that notifies
the user about various battery states and tries to suspend the system when the
battery reaches a critically low level.

Battery Notify continuously observes your system's battery status and level, and
dispatches desktop notifications to keep the user informed about the battery's state.
It is designed to work seamlessly on Linux-based operating systems, leveraging D-Bus
for notifications and system interactions.

Key Features:
  - Real-time battery status monitoring without the overhead of constant polling.
  - Notifications through D-Bus, ensuring compatibility with most Linux desktop environments.
  - Supports automatic system suspension upon critically low battery levels.
  - Recommendations for ideal charge and discharge levels to enhance battery longevity.

Command line flags support is provided for additional functionalities.

Command Line Flags:
  - --help: Print a detailed help message outlining usage and options.
  - --batfile: The path to a battery file in the LFSH.
  - --batmax: Maxmium charge threshold.
  - --batmin: Minimum charge threshold.
  - --dead: Battery percentage level to trigger suspension.

All flags can be invoked with one or two dashes "-"

Usage:
Running the tool is straightforward. In your terminal, simply type:

	$ battery-notify

For a list of available options and their descriptions, use:

	$ battery-notify --help

Examples:

1. Start the battery monitor:

    $ battery-notify

2. Display the help message:

   $ battery-notify --help

The tool will continue running, monitoring the battery status, and will send notifications
as the battery level changes or reaches predefined thresholds.

Future versions might offer more customization options, including setting custom
backlight levels and notification icons.
*/
package main

import (
	"battery-notify/battery"
	"battery-notify/notify"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		// Default provided battery path
		battery_path string

		// Ideal charge threshold for optimizing battery longevity
		ideal_charge_max uint
		ideal_charge_min uint

		// At this level our battery is practically dead
		dead uint
	)

	flag.StringVar(&battery_path, "batfile", "/sys/class/power_supply/BAT0", "Path to the battery file in your system.")
	flag.UintVar(&ideal_charge_max, "batmax", 80, "Max charge level for your battery.")
	flag.UintVar(&ideal_charge_min, "batmin", 20, "Minimum charge level for your battery.")
	flag.UintVar(&dead, "dead", 2, "Battery percentage level to trigger suspension.")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [OPTIONS]\n\n", os.Args[0])
		fmt.Println("OPTIONS:")
		flag.PrintDefaults()
	}
	flag.Parse()

	var notified notify.NotificationStatus = notify.NotNotified	

	battery.MonitorBattery(&notified, battery_path, dead, ideal_charge_min, ideal_charge_max)
}
