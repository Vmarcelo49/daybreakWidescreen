//go:build windows

package main

import (
	"log"

	"github.com/gonutz/w32/v2"
	"github.com/gonutz/wui/v2"
)

func startWindow() {
	window := wui.NewWindow()
	configWindow(window)

	window.Show()
}

func throwErrorMessageWindow(message string) {
	if IsCLIMode {
		log.Panicln(message)
		return
	}
	w32.MessageBox(0, message, "Error", w32.MB_ICONERROR)
}
