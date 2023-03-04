package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

//go:embed style.css
var styleCSS string

func main() {
	app := gtk.NewApplication("com.github.diamondburned.gotk4-examples.gtk4.simple", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	// Load the CSS and apply it globally.
	gtk.StyleContextAddProviderForDisplay(
		gdk.DisplayGetDefault(), loadCSS(styleCSS),
		gtk.STYLE_PROVIDER_PRIORITY_APPLICATION,
	)

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("Port Forwarder")
	window.SetDefaultSize(400, 300)

	// Create a new text view
	tv := gtk.NewTextView()

	txt := createText()

	// Set the text buffer of the text view
	buf := tv.Buffer()
	buf.SetText(txt)

	tv.SetBuffer(buf)

	window.SetChild(tv)

	window.Show()
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

func loadCSS(content string) *gtk.CSSProvider {
	prov := gtk.NewCSSProvider()
	prov.ConnectParsingError(func(sec *gtk.CSSSection, err error) {
		// Optional line parsing routine.
		loc := sec.StartLocation()
		lines := strings.Split(content, "\n")
		log.Printf("CSS error (%v) at line: %q", err, lines[loc.Lines()])
	})
	prov.LoadFromData(content)
	return prov
}
