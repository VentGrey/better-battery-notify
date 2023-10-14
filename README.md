# Battery Notify 🔋📢

Epitaph's battery-notify rewritten in Go + superpowers.

Battery Notify is a comprehensive tool designed to monitor your system's battery status and provide timely desktop notifications. Built with performance and precision in mind, it ensures you're always informed about your device's power state.

This tool is intended to replace another tool in my LeftWM theme, your mileage may vary on other window managers, although this should be generic enough to work in most sensible non ultra-minimalist installations.

## Description 📜
Battery Notify observes the battery's status and level, alerting the user based on predefined thresholds. Key features include:

- Real-time battery status monitoring without constant polling.
- Desktop notifications using D-Bus.
- Automatic system suspension on critically low battery levels.
- Ideal charge and discharge recommendations for prolonged battery life.

## Dependencies 🛠️

- Go 1.19 or later.
- Dbus or a Dbus compatible system.
- A system that supports file watchers.

## Installation & Setup 🚀

You can grab a compiled and optimized binary from the [Releases page](https://github.com/VentGrey/better-battery-notify/releases)

or alternatively you can build it yourself:

Clone the Repository:

```shell
git clone https://github.com/your_username/battery-notify.git
cd battery-notify
```

Build using Makefile:

```shell
make build
```

I won't be adding an `install` target to the makefile because this installation is completely optional. If you are an Epitaph user and want to stick to the provided `battery-manager` I don't blame you, it's more minimalist after all.

Whether you built `battery-notify` yourself or downloaded it make sure to replace your current [Epitaph](https://github.com/VentGrey/Epitaph) binary with the new one.

If you don't use [Epitaph](https://github.com/VentGrey/Epitaph) you can:

Run the Application:

```shell
./battery-notify
```

## Usage 🖥️
Run the application with the command:


```shell
battery-notify
```

To view available options, use:

```shell
battery-notify --help
```

## Supported Platforms 🖥️📱

Battery Notify is designed to work on Linux-based operating systems. I really don't know if this would work on a BSD system or others.

## Contributing 🤝
We welcome contributions! Please see our Contributing Guide for more details.

## License 📄

This project is licensed under the Gnu GPL v3 License - see the LICENSE file for details.

- fsnotify is licensed under the BSD-3-Clause license. All rights reserved to its respective authors.
- dbus is licensed under theBSD-2-Clause license. All rights reserved to its respective authors.

## Support 🌐

For any questions or support, please raise an issue in the GitHub repository or contact the maintainer directly.

---

Made with 💙 by [Your Name/Team Name]
