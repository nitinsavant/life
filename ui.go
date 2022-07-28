package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	minSize = 10
	minGosperSize = 40
)

func promptForInitialState() (selection int, size int) {
	selection = getSelectionFromUser()

	size = getSizeFromUser(selection)

	return selection, size
}

func getSelectionFromUser() (selection int) {
	var err error

	fmt.Println("Welcome to Conway's Game of Life!")
	fmt.Println("\nHow would you like to configure the initial state of the universe?")
	fmt.Println("1) Blinker")
	fmt.Println("2) Toad")
	fmt.Println("3) Beacon")
	fmt.Println("4) Glider")
	fmt.Println("5) Gosper's Glider Gun")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPlease enter a selection between 1-5: ")
		scanner.Scan()
		input := scanner.Text()

		selection, err = strconv.Atoi(input)
		if err != nil { continue }

		if selection < 1 || selection > 5 {
			continue
		}

		break
	}

	fmt.Println(selection)

	return selection
}

func getSizeFromUser(selection int) (size int) {
	var err error

	fmt.Println("\nHow big would you like your universe (in square pixels)?")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if selection == 5 {
			minSize = minGosperSize
		}

		fmt.Printf("\nPlease enter a universe size of at least %v: ", minSize)

		scanner.Scan()
		input := scanner.Text()

		size, err = strconv.Atoi(input)
		if err != nil { continue }

		if selection == 5 && size < minSize {
			fmt.Printf("Gosper's Glider Gun requires a universe size of at least %v.", minSize)
			fmt.Println("")
			continue
		}

		if size < 10 {
			continue
		}

		break
	}

	fmt.Println(size)

	return size
}

