// Copyright 2023 VentGrey. All rights removed.
// Use of this source code is governed by the GNU General Public License v3
// or at your option, any latter version. Such license can be found in the
// LICENSE file. Such license governance does not cover this code dependencies.

/*
Command battery-notify provides a small battery monitoring tool that notifies
the user about various battery states and tries to suspend the system when the
battery reaches a critically low level.

The utility observes the battery's status and level, sending notificcations
to the user, based on certain thresholds. These features are achieved by:

1. Observing battery status and level without polling (using fsnotify).

2. Sending desktop notifications using D-Bus.

3. Suspending the system when the battery reaches a critically low.

Command line flags support is provided for additional functionalities.
*/
package main

import (
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

		// Time in seconds before checking again.
		sleep_time uint
	)

	flag.StringVar(&battery_path, "batfile", "/sys/class/power_supply/BAT0", "Path to the battery file in your system.")
	flag.StringVar(&battery_path, "bf", "/sys/class/power_supply/BAT0", "Path to the battery file in your system [shorthand].")
	flag.UintVar(&ideal_charge_max, "batmax", 80, "Max charge level for your battery.")
	flag.UintVar(&ideal_charge_min, "batmin", 20, "Minimum charge level for your battery.")
	flag.UintVar(&dead, "dead", 2, "Battery percentage level to trigger suspension.")
	flag.UintVar(&sleep_time, "sleep", 10, "Wait time before checking again.")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [OPTIONS]\n\n", os.Args[0])
		fmt.Println("OPTIONS:")
		flag.PrintDefaults()
	}
	flag.Parse()
}
