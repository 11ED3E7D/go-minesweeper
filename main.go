package main

import "fmt"

type Box struct {
	hasBomb           bool
	neighbouringBombs byte
}

var sizeX, sizeY, numOfBombs int16
var mineField [][]Box

func main() {
	fmt.Println("Hello, world")
	sizeX, sizeY = 30, 30
	mineField = initMinefield(sizeX, sizeY)
	printMinefield(mineField)
}

func initMinefield(sizeX, sizeY int16) [][]Box {
	field := make([][]Box, sizeY)
	fmt.Println()
	for i := range field {
		field[i] = make([]Box, sizeX)
		for j := range field[i] {
			field[i][j] = Box{}
		}
	}
	return field
}

func printMinefield(field [][]Box) {
	for i := -2; i < len(field); i++ {
		if i >= 0 {
			fmt.Printf(" %d | ", i)
		} else {
			fmt.Printf("     ")
		}
		for j := 0; j < len(field[0]); j++ {
			switch {
			case i == -2:
				fmt.Printf(" %d ", j)
			case i == -1:
				fmt.Printf(" _ ")
			case i >= 0:
				fmt.Printf(" %d ", field[i][j].neighbouringBombs)
			}
		}
		fmt.Println()
	}
}
