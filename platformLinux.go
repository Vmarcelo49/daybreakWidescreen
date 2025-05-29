//go:build linux

package main

import (
	"fmt"
	"log"
	"os"
)

func throwErrorMessageWindow(message string) {
	log.Panicln(message)
}

func startWindow() {
	fmt.Println("No GUI available for Linux, check -help for CLI options")
	fmt.Println("If you want to use the GUI, please run it on Windows or use Proton")
	os.Exit(0)
}
