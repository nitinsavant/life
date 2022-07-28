package main

import "fmt"

type universe struct {
	selection int
	size int
	state [][]int
	neighborStore [][]int
}

func initializeUniverse(selection int, size int) {
	u.selection = selection
	u.size = size

	u.state = make([][]int, size)
	for i := range u.state {
		u.state[i] = make([]int, size)
	}

	u.neighborStore = make([][]int, size)
	for i := range u.neighborStore {
		u.neighborStore[i] = make([]int, size)
	}

	switch selection {
	case 1:
		u.loadBlinker()
	case 2:
		u.loadToad()
	case 3:
		u.loadBeacon()
	case 4:
		u.loadGlider()
	case 5:
		u.loadGliderGun()
	}
}

func updateUniverse() {
	updateNeighborStore()

	updateState()
}

func updateNeighborStore() {
	numRows := len(u.state)
	numCols := len(u.state[0])

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			cell := cell{col, row}
			u.neighborStore[row][col] = cell.numOfNeighbors()
		}
	}
}

func updateState() {
	numRows := len(u.state)
	numCols := len(u.state[0])

	var num int
	var isAlive bool

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			isAlive = u.state[row][col] == 1
			num = u.neighborStore[row][col]

			if isAlive && (num < 2 || num > 3) {
				u.state[row][col] = 0
			}

			if !isAlive && num == 3 {
				u.state[row][col] = 1
			}
		}
	}
}


func renderUniverse() {
	numRows := len(u.state)
	numCols := len(u.state[0])

	var output string

	for col := 0; col < numCols; col++ {
		for row := 0; row < numRows; row++ {
			if u.state[col][row] == 1 {
				output += aliveMarker
			} else {
				output += deadMarker
			}
		}
		output += "\n"
	}

	// Moves cursor back to row 0, column 0 of an ANSI terminal
	fmt.Printf("\033[0;0H")
	fmt.Printf(output)
}

