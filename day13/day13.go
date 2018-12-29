package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Cart ...struct for location
type Cart struct {
	x                int
	y                int
	currentDirection byte
	nextTurn         string
}

func main() {
	//real inputs
	const xRange = 150
	const yRange = 150
	// const xRange = 8
	// const yRange = 7
	//build 2d array as map
	trackMap := make([][]byte, xRange)
	for i := range trackMap {
		trackMap[i] = make([]byte, yRange)
	}
	fileHandle, _ := os.Open("internetinput.txt")
	// fileHandle, _ := os.Open("day13input.txt")
	// fileHandle, _ := os.Open("day13test2input.txt")
	fileScanner := bufio.NewScanner(fileHandle)
	carts := []Cart{}
	cartCount := 0
	cartIndex := 0
	carts = make([]Cart, cartCount)
	counter := 0
	// fo, err := os.Create("output.txt")
	// check(err)
	//scan in the inputs and get the number of carts
	//there is a better way of handling this by doing a copy/append
	for fileScanner.Scan() {
		trackMap[counter] = []byte(fileScanner.Text())
		cartCount += getCartCount(trackMap[counter])
		counter++
	}
	//output for map
	// for i := range trackMap {
	// 	fmt.Println(string(trackMap[i]))
	// }
	// fmt.Println(string(trackMap[108]))
	//allocate memory for the number of carts
	carts = make([]Cart, cartCount)
	//carts go left, straight, right at intersections
	//maybe change the cartmap position to replace turns with - or |
	for j := range trackMap {
		if bytes.Contains(trackMap[j], []byte("<")) {
			for i := range trackMap[j] {
				if trackMap[j][i] == '<' {
					carts[cartIndex] = Cart{i, j, '<', "left"}
					cartIndex++
				}
			}
		}
		if bytes.Contains(trackMap[j], []byte(">")) {
			for i := range trackMap[j] {
				if trackMap[j][i] == '>' {
					carts[cartIndex] = Cart{i, j, '>', "left"}
					cartIndex++
				}
			}
		}
		if bytes.Contains(trackMap[j], []byte("^")) {
			for i := range trackMap[j] {
				if trackMap[j][i] == '^' {
					carts[cartIndex] = Cart{i, j, '^', "left"}
					cartIndex++
				}
			}
		}
		if bytes.Contains(trackMap[j], []byte("v")) {
			for i := range trackMap[j] {
				if trackMap[j][i] == 'v' {
					carts[cartIndex] = Cart{i, j, 'v', "left"}
					cartIndex++
				}
			}
		}
	}
	if cartCount != cartIndex {
		return
	}
	//carts are correct
	//traverse through the map and get move the carts
	// turn \(92) (x+1), turn /(47)(x-1), straight |(124), intersectin +(43)
	//cart options <(60), >(62), ^(94), v(118)
	//keep a list of the carts at intersections

	// sort.Slice(carts, func(q, r int) bool {
	// 	if carts[q].y == carts[r].y {
	// 		return carts[q].x < carts[r].x
	// 	}
	// 	return carts[q].y < carts[r].y
	// })

	// sort.Slice(carts, func(i, j int) bool {
	// 	switch {
	// 	case carts[i].y < carts[j].y:
	// 		return true
	// 	case carts[i].y > carts[j].y:
	// 		return false
	// 	default:
	// 		return carts[i].x < carts[j].x
	// 	}
	// })

	sort.Slice(carts, func(q, r int) bool {
		if carts[q].x == carts[r].x {
			return carts[q].y < carts[r].y
		}
		return carts[q].x < carts[r].x
	})

	for n := 0; n < 100000000; n++ {
		for i := 0; i < len(carts); i++ {
			// if carts[i].x == 4 && carts[i].y == 109 {
			// 	fmt.Println(string(trackMap[carts[i].y][carts[i].x]))
			// }
			//every cart moves
			if carts[i].currentDirection == '^' {
				carts[i].y--
				// turn \(92) (x+1), turn /(47)(x-1), straight |(124), intersectin +(43)

				if trackMap[carts[i].y][carts[i].x] == 92 || trackMap[carts[i].y][carts[i].x] == byte('r') {
					carts[i].currentDirection = '<'
				} else if trackMap[carts[i].y][carts[i].x] == 47 || trackMap[carts[i].y][carts[i].x] == byte('l') {
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
			} else if carts[i].currentDirection == '>' {
				carts[i].x++
				// turn \(92) (x+1), turn /(47)(x-1), straight |(124), intersectin +(43)
				if trackMap[carts[i].y][carts[i].x] == 92 || trackMap[carts[i].y][carts[i].x] == byte('r') {
					carts[i].currentDirection = 'v'
				} else if trackMap[carts[i].y][carts[i].x] == 47 || trackMap[carts[i].y][carts[i].x] == byte('l') {
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
				if trackMap[carts[i].y][carts[i].x] == 92 || trackMap[carts[i].y][carts[i].x] == byte('r') {
					carts[i].currentDirection = '^'
				} else if trackMap[carts[i].y][carts[i].x] == 47 || trackMap[carts[i].y][carts[i].x] == byte('l') {
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
				}
			} else if carts[i].currentDirection == 'v' {
				carts[i].y++
				// turn \(92) (x+1), turn /(47)(x-1), straight |(124), intersectin +(43)
				if trackMap[carts[i].y][carts[i].x] == 92 || trackMap[carts[i].y][carts[i].x] == byte('r') {
					carts[i].currentDirection = '>'

				} else if trackMap[carts[i].y][carts[i].x] == 47 || trackMap[carts[i].y][carts[i].x] == byte('l') {
					// \
					carts[i].currentDirection = '<'
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
			for j := 0; j < len(carts); j++ {
				if i != j && carts[i].x == carts[j].x && carts[i].y == carts[j].y {
					secondCart := carts[j]
					secCartIndex := j
					fmt.Printf("Removed at x %d and y %d\n", carts[i].x, carts[i].y)
					fmt.Println(carts)
					carts = remove(carts, i)
					for q := 0; q < len(carts); q++ {
						if carts[q] == secondCart {
							secCartIndex = q
						}
					}
					if i > 1 {
						i -= 2
					} else {
						i = 0
					}
					j = 0
					carts = remove(carts, secCartIndex)
				}
			}
		}

		// sort.Slice(carts, func(q, r int) bool {
		// 	if carts[q].y == carts[r].y {
		// 		return carts[q].x < carts[r].x
		// 	}
		// 	return carts[q].y < carts[r].y
		// })

		sort.Slice(carts, func(q, r int) bool {
			if carts[q].x == carts[r].x {
				return carts[q].y < carts[r].y
			}
			return carts[q].x < carts[r].x
		})

		// sort.Slice(carts, func(i, j int) bool {
		// 	switch {
		// 	case carts[i].x < carts[j].x:
		// 		return true
		// 	case carts[i].x > carts[j].x:
		// 		return false
		// 	default:
		// 		return carts[i].y < carts[j].y
		// 	}
		// })

		// for i := 0; i < len(carts); i++ {
		// 	j := i + 1
		// 	for j < len(carts) {
		// 		//For part 1 just print the first collision removal
		// 		if carts[i].x == carts[j].x && carts[i].y == carts[j].y {
		// 			// fmt.Printf("Carts have collided # %d and %d\n", i, j)
		// 			// fmt.Println(carts[i])
		// 			// for q := 0; q < len(carts); q++ {
		// 			// fmt.Printf("x %d and y %d\n", carts[q].x, carts[q].y)
		// 			// }
		// 			secondCart := carts[j]
		// 			secCartIndex := j
		// 			fmt.Printf("Removed at x %d and y %d\n", carts[i].x, carts[i].y)
		// 			fmt.Println(carts)
		// 			carts = remove(carts, i)
		// 			for q := 0; q < len(carts); q++ {
		// 				if carts[q] == secondCart {
		// 					secCartIndex = q
		// 				}
		// 			}
		// 			carts = remove(carts, secCartIndex)
		// 			// fmt.Print("New Carts : ")
		// 			// fmt.Println(carts)
		// 			// fmt.Print("\n\n")
		// 			// loop through again and check for more colisions this turn
		// 			i = 0
		// 			j = 0
		// 			n = 0
		// 			// sort.Slice(carts, func(q, r int) bool { return carts[q].y < carts[r].y })
		// 			sort.Slice(carts, func(q, r int) bool { return carts[q].x < carts[r].x })

		// 			break
		// 		}
		// 		j++
		// 	}
		// }
		// fmt.Printf("Run# %d\n", n+1)
		if len(carts) == 1 {
			break
		}
	}
	fmt.Print(carts[0])
	fmt.Println(carts[0].x, carts[0].y)
	//loop through the carts and check if they intersected
}

func remove(slice []Cart, s int) []Cart {
	return append(slice[:s], slice[s+1:]...)
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
