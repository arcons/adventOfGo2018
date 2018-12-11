package main

import (
	"fmt"
)

func main() {
	fuelGrid := [300][300]int{}
	// square := [9][9]int{}
	//real input is 7689
	serialNumber := 7689
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			//increase x by 1 since we started at 0,0 not 1,1
			rackID := 10 + x
			// start = rackId * y
			start := rackID * y
			powerLevel := start + serialNumber
			powerLevel *= rackID
			hundredDigit := powerLevel / 100 % 10
			powerLevel = hundredDigit - 5
			//save the power level as the position
			fuelGrid[x-1][y-1] = powerLevel
		}
	}

	maxSquareSum := 0
	maxX := 0
	maxY := 0
	squareSum := 0
	// //Part A
	// // loop through and find the largest 3x3 sum
	// for x := 0; x < 300; x++ {
	// 	for y := 0; y < 300; y++ {
	// 		//make sure the square is in bounds
	// 		if x+3 < 300 && y+3 < 300 {
	// 			for ix := 0; ix < 3; ix++ {
	// 				for iy := 0; iy < 3; iy++ {
	// 					squareSum += fuelGrid[x+ix][y+iy]
	// 				}
	// 			}
	// 			// squareSum = fuelGrid[x][y] + fuelGrid[x+1][y] + fuelGrid[x+2][y] + fuelGrid[x+1][y+1] + fuelGrid[x+2][y+1] + fuelGrid[x+2][y+2] + fuelGrid[x][y+1] + fuelGrid[x][y+2]
	// 			if squareSum > maxSquareSum {
	// 				maxSquareSum = squareSum
	// 				maxX = x
	// 				maxY = y
	// 			}
	// 			squareSum = 0
	// 		}
	// 	}
	// }

	//Part B
	//find the largest sum and square size
	maxSquareSize := 1
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			for squareSize := 0; squareSize < 300; squareSize++ {
				if squareSize+x < 300 && squareSize+y < 300 {
					//traverse right
					for ix := 0; ix < squareSize; ix++ {
						for iy := 0; iy < squareSize; iy++ {
							squareSum += fuelGrid[x+ix][y+iy]
						}
					}
					if squareSum > maxSquareSum {
						maxSquareSum = squareSum
						maxX = x
						maxY = y
						maxSquareSize = squareSize
					}
					//reset squareSum to 0
					squareSum = 0
				}
			}
			squareSum = 0
		}
	}
	fmt.Println(maxSquareSum)
	fmt.Println(maxX + 1)
	fmt.Println(maxY + 1)
	fmt.Println(maxSquareSize)
}
