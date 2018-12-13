package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type vector struct {
	posX int
	posY int
	velX int
	velY int
}

//Correct answer ZAEKAJGC
//10577

//
// X:-52761
// Y:-52705
// X:53064
// Y:53074

func main2() {
	fileHandle, _ := os.Open("day10input.txt")
	const inputSize = 327
	fo, err := os.Create("output.txt")
	fileScanner := bufio.NewScanner(fileHandle)
	inputVectors := [inputSize]vector{}
	counter := 0
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	println("Scan in the inputs")
	for fileScanner.Scan() {
		scan := strings.Split(fileScanner.Text(), ">")
		position := strings.Split(strings.Split(scan[0], "<")[1], ",")
		velocity := strings.Split(strings.Split(scan[1], "<")[1], ",")
		// position 1 is x, y
		// print(position[1])
		posX, err := strconv.Atoi(strings.TrimSpace(position[0]))
		check(err)
		posY, err := strconv.Atoi(strings.TrimSpace(position[1]))
		check(err)
		velX, err := strconv.Atoi(strings.TrimSpace(velocity[0]))
		check(err)
		velY, err := strconv.Atoi(strings.TrimSpace(velocity[1]))
		check(err)

		//mark the boundarys
		minX = min(posX, minX)
		maxX = max(posX, maxX)
		minY = min(posY, minY)
		maxY = max(posY, maxY)

		newVector := vector{posX, posY, velX, velY}
		inputVectors[counter] = newVector
		counter++
	}
	// x and y range are the bounds
	xRange := int(math.Abs(float64(minX))) + maxX
	yRange := int(math.Abs(float64(minY))) + maxY
	if xRange%2 == 1 {
		xRange++
	}
	if yRange%2 == 1 {
		yRange++
	}

	//double the size so there is space and we can work solely in positive
	// create a map of the sky
	skyMap := make([][]byte, yRange*2)
	for i := range skyMap {
		skyMap[i] = make([]byte, xRange*2)
	}
	// the center is just the middle of the graph with no velocity
	center := vector{xRange / 2, yRange / 2, 0, 0}
	// fmt.Println(center)
	//22x16, 11,6 is center point
	//print x times
	//center is traditional 0, 0
	for i := 0; i < 4; i++ {
		//loop through and mark all positions then update
		for j := 0; j < inputSize; j++ {
			//mark all of the positions in teh skypMap
			oldVector := inputVectors[j]
			//mark the current spot as an X
			// get the real x and y thats graphable
			graphX := oldVector.posX + center.posX - 1
			graphY := oldVector.posY + center.posY - 1
			// this is the old ne
			skyMap[graphX][graphY] = '.'
			oldVector.posX += oldVector.velX
			oldVector.posY += oldVector.velY

			graphX = oldVector.posX + center.posX - 1
			graphY = oldVector.posY + center.posY - 1
			// println(graphX, graphY)
			skyMap[graphX][graphY] = 'X'
			// not sure if this will pass the reference
			inputVectors[j] = oldVector
		}
		// this prints the dots
		// for x := 0; x < xRange; x++ {
		// 	for y := 0; y < yRange; y++ {
		// 		//replace this with writing the mapped position
		// 		// if the position is marked print the marked position, else print
		// 		if skyMap[x][y] == 'X' {
		// 			fo.Write([]byte(" X"))
		// 		} else {
		// 			fo.Write([]byte(" ."))
		// 		}
		// 	}
		// 	fo.Write([]byte("\n"))
		// }
		fo.Write([]byte("\n\n\n"))
	}
	//write three spaces between input
	fo.Write([]byte("\n\n\n"))
	fo.Write([]byte("Yar"))

	// For more granular writes, open a file for writing.
	check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	defer fo.Close()

}
