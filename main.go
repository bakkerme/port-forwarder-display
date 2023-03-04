package main

import (
	_ "embed"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()

	window := app.NewWindow("Port Forwarder")
	window.Resize(fyne.NewSize(400, 300))

	// Create a new text view
	tv := widget.NewMultiLineEntry()

	txt := createText()

	// Set the text of the text view
	tv.SetText(txt)

	window.SetContent(container.NewVScroll(tv))

	window.ShowAndRun()
}

func createText() string {
	outboundIP := getOutboundIP()
	SSID := getSSID()
	openPorts := getAllOpenPorts()
	sshUsers := getSSHUsers()

	return fmt.Sprintf(
		"Outbound IP: %s\nSSID: %s\nOpen Ports:\n%s\nSSH Users:\n%s",
		outboundIP,
		SSID,
		openPorts,
		sshUsers,
	)
}
