package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

//Cart ...struct for location
type Cart struct {
	x                int
	y                int
	currentDirection byte
	nextTurn         string
}

func main() {
	// fmt.Print("Hello")
	//real inputs
	// const xRange = 151
	// const yRange = 150
	const xRange = 14
	const yRange = 6
	//build 2d array as map
	trackMap := make([][]byte, yRange)
	for i := range trackMap {
		trackMap[i] = make([]byte, xRange)
	}

	fileHandle, _ := os.Open("day13input.txt")
	fileScanner := bufio.NewScanner(fileHandle)
	carts := []Cart{}
	cartCount := 0
	cartIndex := 0
	carts = make([]Cart, cartCount)
	counter := 0
	//scan in the inputs and get the number of carts
	//there is a better way of handling this by doing a copy/append
	for fileScanner.Scan() {
		trackMap[counter] = []byte(fileScanner.Text())
		cartCount += getCartCount(trackMap[counter])
		counter++
	}
	//allocate memory for the number of carts
	carts = make([]Cart, cartCount)
	//carts go left, straight, right at intersections
	//maybe change the cartmap position to replace turns with - or |
	for i := range trackMap {
		fmt.Println(string(trackMap[i]))
	}
	for j := range trackMap {
		if bytes.Contains(trackMap[j], []byte("<")) {
			for i := range trackMap[j] {
				if trackMap[j][i] == '<' {
					carts[cartIndex] = Cart{i, j, '<', "left", false}
					cartIndex++
					trackMap[j][i] = '-'
				}
			}
		}
		if bytes.Contains(trackMap[j], []byte(">")) {
			for i := range trackMap[j] {
				if trackMap[j][i] == '>' {
					carts[cartIndex] = Cart{i, j, '>', "left", false}
					cartIndex++
					trackMap[j][i] = '-'
				}
			}
		}
		if bytes.Contains(trackMap[j], []byte("^")) {
			for i := range trackMap[j] {
				if trackMap[j][i] == '^' {
					carts[cartIndex] = Cart{i, j, '^', "left", false}
					cartIndex++
					trackMap[j][i] = '|'

				}
			}
		}
		if bytes.Contains(trackMap[j], []byte("v")) {
			for i := range trackMap[j] {
				if trackMap[j][i] == 'v' {
					// fmt.Println(j)
					carts[cartIndex] = Cart{i, j, 'v', "left", false}
					cartIndex++
					trackMap[j][i] = '|'
				}
			}
		}
	}
	collidingX := 0
	collidingY := 0

	//carts are correct
	//traverse through the map and get move the carts
	// turn \(92) (x+1), turn /(47)(x-1), straight |(124), intersectin +(43)
	//cart options <(60), >(62), ^(94), v(118)
	//keep a list of the carts at intersections
	for n := 0; n < 100000000; n++ {
		for i := range carts {
			if carts[i].currentDirection == '>' {
				carts[i].x++
				// turn \(92) (x+1), turn /(47)(x-1), straight |(124), intersectin +(43)
				if trackMap[carts[i].y][carts[i].x] == 92 {
					carts[i].currentDirection = 'v'
				} else if trackMap[carts[i].y][carts[i].x] == 47 {
					carts[i].currentDirection = '^'
				} else if trackMap[carts[i].y][carts[i].x] == 43 {
					if carts[i].nextTurn == "left" {
						carts[i].nextTurn = "str"
						carts[i].currentDirection = '^'
					} else if carts[i].nextTurn == "str" {
						carts[i].nextTurn = "right"
						carts[i].currentDirection = '>'
					} else if carts[i].nextTurn == "right" {
						carts[i].nextTurn = "left"
						carts[i].currentDirection = 'v'
					}
				}
			} else if carts[i].currentDirection == '<' {
				carts[i].x--
				// turn \(92) (x+1), turn /(47)(x-1), straight |(124), intersectin +(43)
				if trackMap[carts[i].y][carts[i].x] == 92 {
					carts[i].currentDirection = '^'
				} else if trackMap[carts[i].y][carts[i].x] == 47 {
					carts[i].currentDirection = 'v'
				} else if trackMap[carts[i].y][carts[i].x] == 43 {
					if carts[i].nextTurn == "left" {
						carts[i].nextTurn = "str"
						carts[i].currentDirection = 'v'
					} else if carts[i].nextTurn == "str" {
						carts[i].nextTurn = "right"
						carts[i].currentDirection = '<'
					} else if carts[i].nextTurn == "right" {
						carts[i].nextTurn = "left"
						carts[i].currentDirection = '^'
					}
					carts[i].onIntersection = true
				} else {
					carts[i].onIntersection = false
				}
			} else if carts[i].currentDirection == '^' {
				// left turn
				carts[i].y--
				// turn \(92) (x+1), turn /(47)(x-1), straight |(124), intersectin +(43)

				if trackMap[carts[i].y][carts[i].x] == 92 {
					carts[i].currentDirection = '<'
				} else if trackMap[carts[i].y][carts[i].x] == 47 {
					// right
					carts[i].currentDirection = '>'
				} else if trackMap[carts[i].y][carts[i].x] == 43 {
					// +
					if carts[i].nextTurn == "left" {
						carts[i].nextTurn = "str"
						carts[i].currentDirection = '<'
					} else if carts[i].nextTurn == "str" {
						carts[i].nextTurn = "right"
						carts[i].currentDirection = '^'
					} else if carts[i].nextTurn == "right" {
						carts[i].nextTurn = "left"
						carts[i].currentDirection = '>'
					}
				}
			} else if carts[i].currentDirection == 'v' {
				carts[i].y++
				// turn \(92) (x+1), turn /(47)(x-1), straight |(124), intersectin +(43)
				if trackMap[carts[i].y][carts[i].x] == 92 {
					carts[i].currentDirection = '>'

				} else if trackMap[carts[i].y][carts[i].x] == 47 {
					// \
					carts[i].currentDirection = '<'
					carts[i].onIntersection = false
				} else if trackMap[carts[i].y][carts[i].x] == 43 {
					// +
					if carts[i].nextTurn == "left" {
						carts[i].nextTurn = "str"
						carts[i].currentDirection = '>'
					} else if carts[i].nextTurn == "str" {
						carts[i].nextTurn = "right"
						carts[i].currentDirection = 'v'
					} else if carts[i].nextTurn == "right" {
						carts[i].nextTurn = "left"
						carts[i].currentDirection = '<'
					}
				}
			}
		}

		for i := 0; i < len(carts); i++ {
			j := i + 1
			for j < len(carts) {
				if carts[i].x == carts[j].x && carts[i].y == carts[j].y {
					collidingX = carts[i].x
					collidingY = carts[i].y
					break
				}
				j++
			}
			if collidingX != 0 {
				break
			}
		}
		if collidingX != 0 {
			break
		}
	}

	fmt.Println(collidingX, collidingY)
	//loop through the carts and check if they intersected
}

func getNextLocation(char byte, x int, y int) (int, int) {
	//left turn
	if char == 92 {

	} else if char == 47 {

	}
	return 3, 7
}

func getCartCount(row []byte) int {
	cartCount := 0
	if bytes.Contains(row, []byte("<")) {
		for i := range row {
			if row[i] == '<' {
				cartCount++
			}
		}
	}
	if bytes.Contains(row, []byte(">")) {
		for i := range row {
			if row[i] == '>' {
				cartCount++
			}
		}
	}
	if bytes.Contains(row, []byte("^")) {
		for i := range row {
			if row[i] == '^' {
				cartCount++
			}
		}
	}
	if bytes.Contains(row, []byte("v")) {
		for i := range row {
			if row[i] == 'v' {
				cartCount++
			}
		}
	}
	return cartCount
}
