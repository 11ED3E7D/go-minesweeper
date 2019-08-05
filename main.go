package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Box struct {
	hasBomb           bool
	neighbouringBombs byte
	value             byte
}

var sizeX, sizeY, numOfBombs, threshold int
var mineField [][]Box

func main() {
	gameOver := false
	quitGame := false
	fmt.Println("Hello, world")
	sizeX, sizeY, threshold = 25, 25, 10
	mineField = initMinefield(sizeX, sizeY)
	fmt.Println()
	printMinefield(mineField)
	for !quitGame {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Command > ")
		input, _ := reader.ReadString('\n')
		switch os := runtime.GOOS; os {
		case "windows":
			input = strings.Replace(input, "\r\n", "", -1)
		default:
			input = strings.Replace(input, "\n", "", -1)
		}
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
			quitGame = true
		case "peek":
			if len(args) == 2 && !gameOver {
				y, err0 := strconv.ParseInt(args[0], 10, 64)
				if err0 == nil && int(y) < sizeY {
					x, err1 := strconv.ParseInt(args[1], 10, 64)
					if err1 == nil && int(x) < sizeX {
						gameOver = peek(mineField, int(y), int(x)) // y and then x
					}
				}
			}
			printMinefield(mineField)
		case "mark":
			if len(args) == 2 && !gameOver {
				y, err0 := strconv.ParseInt(args[0], 10, 64)
				if err0 == nil && int(y) < sizeY {
					x, err1 := strconv.ParseInt(args[1], 10, 64)
					if err1 == nil && int(x) < sizeX {
						if mineField[y][x].value == '#' {
							mineField[y][x].value = 'O'
						} else {
							mineField[y][x].value = '#'
						}
					}
				}
			}
			printMinefield(mineField)
		case "peekAll":
			for i := range mineField {
				for j := range mineField[i] {
					peek(mineField, int(i), int(j))
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
			gameOver = false
			fmt.Println(" =============== New Game =============== ")
		case "size":
			if len(args) == 2 {
				y, err0 := strconv.ParseInt(args[0], 10, 64)
				if err0 == nil && y > 0 {
					x, err1 := strconv.ParseInt(args[1], 10, 64)
					if err1 == nil && x > 0 {
						sizeX = int(x)
						sizeY = int(y)
						mineField = initMinefield(sizeX, sizeY)
						printMinefield(mineField)
						fmt.Println(" =============== New Game =============== ")
						gameOver = false
					}
				}
			}
		case "threshold":
			if len(args) == 1 {
				num, err := strconv.ParseInt(args[0], 10, 64)
				if err == nil && num >= 0 && num <= 100 {
					threshold = int(num)
				}
			}
		default:
			printHelp()
		}
		if gameOver {
			fmt.Println(" =============== Game Over =============== ")
		} else if checkWin(mineField) {
			fmt.Println(" =============== Winner! ==============")
		}
	}
}

func checkWin(field [][]Box) bool {
	win := true
	for i := range field {
		for j := range field[i] {
			if field[i][j].hasBomb && field[i][j].value != '#' {
				return false
			}
		}
	}
	return win
}

// Flood-fill (node, target-color, replacement-color):
//  1. If target-color is equal to replacement-color, return.
//  2. ElseIf the color of node is not equal to target-color, return.
//  3. Else Set the color of node to replacement-color.
//  4. Perform Flood-fill (one step to the south of node, target-color, replacement-color).
//     Perform Flood-fill (one step to the north of node, target-color, replacement-color).
//     Perform Flood-fill (one step to the west of node, target-color, replacement-color).
//     Perform Flood-fill (one step to the east of node, target-color, replacement-color).
//  5. Return.
func peek(field [][]Box, y int, x int) bool {
	if field[y][x].value == ' ' {
		return false
	} else if field[y][x].hasBomb {
		field[y][x].value = 'X'
		return true
	} else if field[y][x].neighbouringBombs > 0 {
		// Add 48 to get ascii number. This is fine as neighbouring bombs is <= 8
		field[y][x].value = field[y][x].neighbouringBombs + 48
	} else {
		field[y][x].value = ' '
		if y+1 < len(field) {
			peek(field, y+1, x)
		}
		if y-1 >= 0 {
			peek(field, y-1, x)
		}
		if x-1 >= 0 {
			peek(field, y, x-1)
		}
		if x+1 < len(field[0]) {
			peek(field, y, x+1)
		}
	}
	return false
}

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("  help - show this menu")
	fmt.Println("  quit - quit the game")
	fmt.Println("  peek <y> <x> - uncovers the box at that position")
	fmt.Println("  mark <y> <x> - mark position as bomb")
	fmt.Println("  peekAll - uncover all boxes")
	fmt.Println("  hideAll - covers all boxes")
	fmt.Println("  new  - creates a new minefield")
	fmt.Println("  size <height> <width> - Sets the new size of field and creates new game")
	fmt.Println("  threshold <num: 0 - 100> - Changes the probability of bombs being placed")
}

func initMinefield(sizeX, sizeY int) [][]Box {
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

// Used to set box's neighbouring bomb counters
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
	threshold := threshold
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
