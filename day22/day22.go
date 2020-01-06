package main

import (
	"fmt"
)

func main() {
	depth := 7863
	inittargetx := 14
	inittargety := 760
	// depth: 7863
	// target: 14,760
	// depth := 510
	// inittargetx := 10
	// inittargety := 10
	targetx := inittargetx + 1
	targety := inittargety + 1

	caveMap := make([][]byte, targetx)
	for i := range caveMap {
		caveMap[i] = make([]byte, targety)
	}

	caveMapErro := make([][]int, targetx)
	for i := range caveMapErro {
		caveMapErro[i] = make([]int, targety)
	}

	caveMap[inittargetx][inittargety] = 'T'
	for y := 0; y < targety; y++ {
		er := getErro(depth, getGeologic(0, y, depth, caveMapErro))
		caveMapErro[0][y] = er
	}

	for x := 0; x < targetx; x++ {
		er := getErro(depth, getGeologic(x, 0, depth, caveMapErro))
		caveMapErro[x][0] = er
	}

	for y := 0; y < targety; y++ {
		for x := 0; x < targetx; x++ {
			er := getErro(depth, getGeologic(x, y, depth, caveMapErro))
			caveMapErro[x][y] = er
			fmt.Println(x, y)
		}
	}

	// A region's erosion level is its geologic index plus the cave system's depth, all modulo 20183
	// 	If the erosion level modulo 3 is 0, the region's type is rocky.
	// If the erosion level modulo 3 is 1, the region's type is wet.
	// If the erosion level modulo 3 is 2, the region's type is narrow.
	fmt.Println(depth)

	for y := 0; y < targety; y++ {
		for x := 0; x < targetx; x++ {
			if caveMap[x][y] != 'T' {
				caveMap[x][y] = setRegion(caveMap, depth, caveMapErro, x, y)
			}
		}
	}
	risk := 0
	// 0 for rocky regions, 1 for wet = regions, and 2 for narrow |
	for y := 0; y < inittargety+1; y++ {
		for x := 0; x < inittargetx+1; x++ {
			if caveMap[x][y] == '=' {
				risk++
			}
			if caveMap[x][y] == '|' {
				risk += 2
			}
		}
	}
	fmt.Println(risk)
}

func setRegion(caveMap [][]byte, depth int, caveMapErro [][]int, x int, y int) byte {
	er := caveMapErro[x][y]
	regionType := byte('.')
	if er%3 == 0 {
		regionType = byte('.')
	}
	if er%3 == 1 {
		regionType = byte('=')
	}
	if er%3 == 2 {
		regionType = byte('|')
	}
	return regionType
}

// 	The region at 0,0 (the mouth of the cave) has a geologic index of 0.
// The region at the coordinates of the target has a geologic index of 0.
// If the region's Y coordinate is 0, the geologic index is its X coordinate times 16807.
// If the region's X coordinate is 0, the geologic index is its Y coordinate times 48271.
// Otherwise, the region's geologic index is the result of multiplying the erosion levels of the regions at X-1,Y and X,Y-1.
func getGeologic(x int, y int, depth int, caveMapErro [][]int) int {
	if x == 0 && y == 0 {
		return 0
	}
	if x == 14 && y == 760 {
		return 0
	}
	if y == 0 {
		return x * 16807
	}
	if x == 0 {
		return y * 48271
	}
	if caveMapErro[x-1][y] == 0 || caveMapErro[x][y-1] == 0 {
		return caveMapErro[x-1][y] * caveMapErro[x][y-1]
	}
	return getErro(depth, getGeologic(x-1, y, depth, caveMapErro)) * getErro(depth, getGeologic(x, y-1, depth, caveMapErro))
}

func getErro(depth int, geo int) int {
	return (geo + depth) % 20183
}
