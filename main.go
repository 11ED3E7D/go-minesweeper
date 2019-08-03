package main

import (
	"fmt"
	"strings"
)

type Box struct {
	hasBomb           bool
	neighbouringBombs byte
	value             byte
}

var sizeX, sizeY, numOfBombs int16
var mineField [][]Box

func main() {
	fmt.Println("Hello, world")
	sizeX, sizeY = 25, 25
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
	// yLen is the padding for the numbers on the y axis
	yLen := numOfDigits(len(field) + 1)
	for i := -2; i < len(field); i++ {
		if i >= 0 {
			fmt.Printf("%v%d | ", strings.Repeat(" ", yLen-numOfDigits(i)), i)
		} else {
			fmt.Printf("%v   ", strings.Repeat(" ", yLen))
		}
		// xLen is the padding for the x numbers and each Box cell
		// Have to use 0 instead of i as -2 < i < n
		xLen := numOfDigits(len(field[0]) + 1)
		for j := 0; j < len(field[0]); j++ {
			switch {
			case i == -2:
				fmt.Printf(" %v%d ", strings.Repeat(" ", xLen-numOfDigits(j)), j)
			case i == -1:
				fmt.Printf("%v_ ", strings.Repeat(" ", xLen))
			case i >= 0:
				fmt.Printf(" %v%d ", strings.Repeat(" ", xLen-numOfDigits(int(field[i][j].value))), field[i][j].neighbouringBombs)
			}
		}
		fmt.Println()
	}
}

func numOfDigits(num int) int {
	digits := 0
	if num == 0 {
		digits++
	}
	for num != 0 {
		num = num / 10
		digits++
	}
	return digits
}
