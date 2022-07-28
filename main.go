package main

import (
	"time"
)

const (
	aliveMarker     = "o "
	deadMarker      = "  "
	delay = 200
)

var u universe

func main() {
	selection, size := promptForInitialState()

	initializeUniverse(selection, size)

	clearConsole()

	runGame()
}

func runGame() {
	for {
		renderUniverse()
		updateUniverse()
		time.Sleep(delay * time.Millisecond)
	}
}

