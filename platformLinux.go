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
	fmt.Println("No GUI available for Linux")
	os.Exit(1)
}
