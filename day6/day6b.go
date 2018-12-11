package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Print("Dyay 6 AOCC")
	fileHandle, _ := os.Open("day6input.txt")
	var input = [50]coordinate{}
	var counter = 0
	fileScanner := bufio.NewScanner(fileHandle)
	//definitely a clean way of doing this
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		xy := strings.Split(fileScanner.Text(), ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(strings.TrimSpace(xy[1]))
		input[counter] = coordinate{x, y, fileScanner.Text()}
		counter++
	}
	//go to area +1 for boundary
	area := [353][360]taxiMap{}
	// mark each coordinate with its index
	var mostValid = 0
	for i := 0; i < 352; i++ {
		for j := 0; j < 359; j++ {
			area[i][j] = taxiMap{"X", 10000}
		}
	}
	//build map
	for i := 0; i < 50; i++ {
		xCoord := input[i].x
		yCoord := input[i].y
		//unmarked = 500
		//mark the square
		area[xCoord][yCoord].minStartID = strconv.Itoa(i)
		//convert all other squares to the value
		for j := 0; j < 352; j++ {
			//mark all the taxi distances from this point
			//taxi distance
			for k := 0; k < 359; k++ {
				//non a mark position
				// find the taxi distance to the point j, k
				// sqrt(x1-x2
				taxiDistance := math.Abs(float64(xCoord)-float64(j)) + math.Abs(float64(yCoord)-float64(k))
				//check if the distance is the same
				if area[j][k].distanceFromStart == taxiDistance {
					area[j][k].minStartID = "N"
				} else if area[j][k].distanceFromStart > taxiDistance {
					area[j][k].distanceFromStart = taxiDistance
					area[j][k].minStartID = strconv.Itoa(i)
					//get the start position and calulate until that point is hit
				}
			}
		}
		//mark all of the squares around it
	}
	//figure out how to check for the valid input
	invalidID := []int{}
	//go around edges and store the ids that are on the edge, if they exist on the edge they are invalid
	for i := 0; i < 352; i++ {
		//top row
		//errors out if theres an N and doesnt perform the task
		number, _ := strconv.Atoi(area[i][0].minStartID)
		if !stringInSlice(number, invalidID) && number != 0 {
			invalidID = append(invalidID, number)
		}
		//bot row
		number, _ = strconv.Atoi(area[i][358].minStartID)
		if !stringInSlice(number, invalidID) && number != 0 {
			invalidID = append(invalidID, number)
		}
	}
	//y coordinate, mislabeled my y coordinate you goofed here and it broke everything
	for i := 0; i < 359; i++ {
		number, _ := strconv.Atoi(area[0][i].minStartID)
		if !stringInSlice(number, invalidID) && number != 0 {
			invalidID = append(invalidID, number)
		}
		number, _ = strconv.Atoi(area[351][i].minStartID)
		if !stringInSlice(number, invalidID) && number != 0 {
			invalidID = append(invalidID, number)
		}
	}
	//find the idea with the largest count
	validSquare := [50]int{}
	for n := 0; n < 50; n++ {
		for i := 0; i < 352; i++ {
			for j := 0; j < 358; j++ {
				//if n == N then its invalid
				if area[i][j].minStartID != "N" && area[i][j].minStartID == strconv.Itoa(n) && !stringInSlice(n, invalidID) {
					validSquare[n]++
				}
			}
		}
	}
	//check the x's
	fmt.Print("\n")
	fmt.Print("Invalid IDs")
	fmt.Print(invalidID)
	fmt.Print("\n")
	for i := 0; i < 50; i++ {
		//Part A Solution
		if !stringInSlice(i, invalidID) && (validSquare[i] > mostValid) {
			mostValid = validSquare[i]
			fmt.Print(" \n")

		}
	}

	fmt.Print("\n")
	fmt.Print("Answer ")
	fmt.Print(mostValid)
}