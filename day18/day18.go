package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Gday m8")
	//create 2d array bytemap start 10x10 for the test
	// treeMap := make([][]byte, 50)
	// for i := range treeMap {
	// 	treeMap[i] = make([]byte, 50)
	// }
	//open ground (.)46, trees (|)124, or a lumberyard (#)35
	//create 2d array bytemap start 10x10 for the test
	treeMap := make([][]byte, 10)
	for i := range treeMap {
		treeMap[i] = make([]byte, 10)
	}

	//make a temp since all of the updates are done at once and are going to be passed into the temp one
	tempTreeMap := make([][]byte, 10)
	for i := range tempTreeMap {
		tempTreeMap[i] = make([]byte, 10)
	}
	//populate the treemap
	fileHandle, _ := os.Open("day18testinput.txt")
	fileScanner := bufio.NewScanner(fileHandle)
	counter := 0
	for fileScanner.Scan() {
		treeMap[counter] = []byte(fileScanner.Text())
		counter++
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(string(treeMap[j]))
		}
		fmt.Println()
		fmt.Println()
		//clear the temptree map everytime
		tempTreeMap := make([][]byte, 10)
		for i := range tempTreeMap {
			tempTreeMap[i] = make([]byte, 10)
		}
		// tempTreeMap = treeMap
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				tempTreeMap[j][k] = checkCell(treeMap, j, k)
				// fmt.Println(k, j)
			}
		}
		treeMap = tempTreeMap
	}

	numLum := 0
	numTree := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if treeMap[i][j] == '|' {
				numTree++
			} else if treeMap[i][j] == '#' {
				numLum++
			}
		}
	}
	fmt.Println(numTree, numLum)
	fmt.Println(numLum * numTree)
}

// An open acre will become filled with trees if three or more adjacent acres contained trees. Otherwise, nothing happens.
// An acre filled with trees will become a lumberyard if three or more adjacent acres were lumberyards. Otherwise, nothing happens.
// An acre containing a lumberyard will remain a lumberyard if it was adjacent to at least one other lumberyard and at least one acre containing trees. Otherwise, it becomes open.
//missing acres arn't counted
func checkCell(treeMap [][]byte, x int, y int) byte {
	//get surrounding trees
	//get the spaces above current location
	// 0 1 2
	// 3 X 4
	// 5 6 7
	currentByte := treeMap[x][y]
	neighbors := [8]byte{}
	//no edge cases this is super gross can definitely be quicker with a swithc
	//error is assignment logic
	if (x-1) < 0 || (x+1) > 9 || (y-1) < 0 || (y+1) > 9 {
		if (x - 1) < 0 {
			if (y - 1) < 0 {
				neighbors[4] = treeMap[x+1][y]
				neighbors[6] = treeMap[x][y+1]
				neighbors[7] = treeMap[x+1][y+1]
			} else if (y + 1) > 9 {
				neighbors[1] = treeMap[x][y-1]
				neighbors[2] = treeMap[x+1][y-1]
				neighbors[4] = treeMap[x+1][y]
			} else {
				neighbors[1] = treeMap[x][y-1]
				neighbors[2] = treeMap[x+1][y-1]
				neighbors[4] = treeMap[x+1][y]
				neighbors[6] = treeMap[x][y+1]
				neighbors[7] = treeMap[x+1][y+1]
			}
		}
		if (x + 1) > 9 {
			if (y - 1) < 0 {
				neighbors[5] = treeMap[x-1][y+1]
				neighbors[6] = treeMap[x][y+1]
				neighbors[3] = treeMap[x-1][y]
			} else if (y + 1) > 9 {
				neighbors[3] = treeMap[x-1][y]
				neighbors[0] = treeMap[x-1][y-1]
				neighbors[1] = treeMap[x][y-1]
			} else {
				neighbors[0] = treeMap[x-1][y-1]
				neighbors[1] = treeMap[x][y-1]
				neighbors[3] = treeMap[x-1][y]
				neighbors[5] = treeMap[x-1][y+1]
				neighbors[6] = treeMap[x][y+1]
			}
		}
		// case would be better here but I dont have internet lol
		if (x + 1) > 9 {
			if (y - 1) < 0 {
				neighbors[5] = treeMap[x-1][y+1]
				neighbors[6] = treeMap[x][y+1]
				neighbors[3] = treeMap[x-1][y]
			} else if (y + 1) > 9 {
				neighbors[3] = treeMap[x-1][y]
				neighbors[0] = treeMap[x-1][y-1]
				neighbors[1] = treeMap[x][y-1]
			} else {
				neighbors[0] = treeMap[x-1][y-1]
				neighbors[1] = treeMap[x][y-1]
				neighbors[3] = treeMap[x-1][y]
				neighbors[5] = treeMap[x-1][y+1]
				neighbors[6] = treeMap[x][y+1]
			}
		}
		//botton row
		if (y - 1) < 0 {
			if (x - 1) < 0 {
				neighbors[4] = treeMap[x+1][y]
				neighbors[6] = treeMap[x][y+1]
				neighbors[7] = treeMap[x+1][y+1]
			} else if (x + 1) > 9 {
				neighbors[3] = treeMap[x-1][y]
				neighbors[5] = treeMap[x-1][y+1]
				neighbors[6] = treeMap[x][y+1]
			} else {
				neighbors[3] = treeMap[x-1][y]
				neighbors[4] = treeMap[x+1][y]
				neighbors[5] = treeMap[x-1][y+1]
				neighbors[6] = treeMap[x][y+1]
				neighbors[7] = treeMap[x+1][y+1]
			}
		}
		if (y + 1) > 9 {
			if (x - 1) < 0 {
				neighbors[1] = treeMap[x][y-1]
				neighbors[2] = treeMap[x+1][y-1]
				neighbors[4] = treeMap[x+1][y]
			} else if (x + 1) > 9 {
				neighbors[0] = treeMap[x-1][y-1]
				neighbors[1] = treeMap[x][y-1]
				neighbors[3] = treeMap[x-1][y]
			} else {
				neighbors[0] = treeMap[x-1][y-1]
				neighbors[1] = treeMap[x][y-1]
				neighbors[2] = treeMap[x+1][y-1]
				neighbors[3] = treeMap[x-1][y]
				neighbors[4] = treeMap[x+1][y]
			}
		}
	} else {
		//non edge on anyside
		neighbors[0] = treeMap[x-1][y-1]
		neighbors[1] = treeMap[x][y-1]
		neighbors[2] = treeMap[x+1][y-1]
		neighbors[3] = treeMap[x-1][y]
		neighbors[4] = treeMap[x+1][y]
		neighbors[5] = treeMap[x-1][y+1]
		neighbors[6] = treeMap[x][y+1]
		neighbors[7] = treeMap[x+1][y+1]
	}
	neighborEmpty := 0
	neighborTrees := 0
	neighborLumbr := 0
	for i := range neighbors {
		if neighbors[i] == '.' {
			neighborEmpty++
		}
		if neighbors[i] == '|' {
			neighborTrees++
		}
		if neighbors[i] == '#' {
			neighborLumbr++
		}
	}
	// 0 1 2
	// 3 X 4
	// 5 6 7
	//open ground (.)46, trees (|)124, or a lumberyard (#)35
	// if x == 7 {
	// 	fmt.Print("Err here")
	// }
	if currentByte == '.' && neighborTrees > 2 {
		return '|'
	}
	if currentByte == '|' && neighborLumbr > 2 {
		return '#'
	}
	if currentByte == '#' {
		if neighborLumbr > 0 && neighborTrees > 0 {
			return '#'
		}
		return '.'
	}

	return currentByte
}
