package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Box struct {
	hasBomb           bool
	neighbouringBombs byte
	value             string
}

var sizeX, sizeY, numOfBombs int16
var mineField [][]Box

func main() {
	fmt.Println("Hello, world")
	sizeX, sizeY = 25, 25
	mineField = initMinefield(sizeX, sizeY)
	fmt.Println()
	printMinefield(mineField)
}

func initMinefield(sizeX, sizeY int16) [][]Box {
	field := make([][]Box, sizeY)
	// Decide where all the bombs will be
	for i := range field {
		field[i] = make([]Box, sizeX)
		for j := range field[i] {
			field[i][j] = genBox()
		}
	}
	// Determine neighbouring bombs
	for i := range field {
		for j := range field[i] {
			checkSurroundings(field, i, j, &(field[i][j]))
		}
	}
	return field
}

func checkSurroundings(field [][]Box, y, x int, box *Box) {
	var bombCount byte = 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Ensure to skip own box and outside edges
			if !(i == 0 && j == 0) && ((x+j) >= 0 && (y+i) >= 0) && ((x+j) < len(field[0]) && (y+i) < len(field)) {
				if field[y+i][x+j].hasBomb {
					bombCount++
				}
			}
		}
	}
	box.neighbouringBombs = bombCount
	if box.hasBomb {
		box.value = `X`
	} else {
		box.value = fmt.Sprintf("%d", bombCount)
	}
}

func genBox() Box {
	if determineBomb() {
		return Box{hasBomb: true}
	} else {
		return Box{}
	}
}

func determineBomb() bool {
	min := 0
	max := 100
	threshold := 30
	return (rand.Intn(max-min) + min) < threshold
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
				fmt.Printf(" %v%v ", strings.Repeat(" ", xLen-1), field[i][j].value)
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
