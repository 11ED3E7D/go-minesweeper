package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Box struct {
	hasBomb           bool
	neighbouringBombs byte
	value             byte
}

var sizeX, sizeY, numOfBombs uint64
var mineField [][]Box

func main() {
	gameOver := false
	fmt.Println("Hello, world")
	sizeX, sizeY = 25, 25
	mineField = initMinefield(sizeX, sizeY)
	fmt.Println()
	printMinefield(mineField)
	for !gameOver {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Command > ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\r\n", "", -1)
		args := strings.Split(input, " ")
		if len(args) > 0 {
			input = args[0]
			args = args[1:]
		}
		switch input {
		case "help":
			printHelp()
		case "quit":
			fmt.Println("quiting game...")
			gameOver = true
		case "peek":
			if len(args) == 2 {
				y, err0 := strconv.ParseUint(args[0], 10, 64)
				if err0 == nil && y < (sizeY-1) {
					x, err1 := strconv.ParseUint(args[1], 10, 64)
					if err1 == nil && x < (sizeX-1) {
						peek(mineField, y, x) // y and then x
					}
				}
			}
			printMinefield(mineField)
		case "peekAll":
			for i := range mineField {
				for j := range mineField[i] {
					peek(mineField, uint64(i), uint64(j))
				}
			}
			printMinefield(mineField)
		case "hideAll":
			for i := range mineField {
				for j := range mineField[i] {
					mineField[i][j].value = 'O'
				}
			}
			printMinefield(mineField)
		case "new":
			mineField = initMinefield(sizeX, sizeY)
			printMinefield(mineField)
		default:
			printHelp()
		}
	}
}

func peek(field [][]Box, y uint64, x uint64) {
	if field[y][x].hasBomb {
		field[y][x].value = 'X'
	} else if field[y][x].neighbouringBombs == 0 {
		field[y][x].value = ' '
	} else {
		field[y][x].value = field[y][x].neighbouringBombs + 48
	}
}

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("  help - show this menu")
	fmt.Println("  quit - quit the game")
	fmt.Println("  peek <y> <x> - uncovers the box at that location")
	fmt.Println("  peekAll - uncover all boxes")
	fmt.Println("  hideAll - covers all boxes")
	fmt.Println("  new  - creates a new minefield")
}

func initMinefield(sizeX, sizeY uint64) [][]Box {
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
}

func genBox() Box {
	if determineBomb() {
		return Box{true, 0, 'O'}
	} else {
		return Box{false, 0, 'O'}
	}
}

func determineBomb() bool {
	min := 0
	max := 100
	threshold := 10
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
				fmt.Printf(" %v%c ", strings.Repeat(" ", xLen-1), field[i][j].value)
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
