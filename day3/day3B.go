package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	input := day3input()
	//find number of overlaps
	//build out giant map of 1000x1000
	fabric := [10000][10000]int{}
	overlaps := [1295]bool{}
	var intersect = 0
	for _, s := range input {
		// split the given string into characters
		//use spaces to delimit since triming is weird with t prefix and suffix
		splitSpace := strings.Split(s, " ");
		// claimId := strings.Trim(splitSpace[0], "#")
		// leftRight at 0, topBottom at 1
		edge := strings.Split(splitSpace[2], ",")
		widthHeight := strings.Split(splitSpace[3], "x")
		//width at 0, height at 1
		leftRight, err := strconv.Atoi(edge[0])
		if err != nil {
			// handle error
		 }
		// fmt.Print(leftRight);
		topBottom, err := strconv.Atoi(strings.Replace(edge[1], ":", "", -1))
		if err != nil {
			// handle error
		 }
		width, err := strconv.Atoi(widthHeight[0])
		if err != nil {
			// handle error
		 }
		height, err := strconv.Atoi(widthHeight[1])
		if err != nil {
			// handle error
		 }
		//start at shift
		//find the fabric that never intersects
		for x := 0; x < width; x++ {
			for y:= 0; y < height; y++ {
				// mark all items in the square with the shift
				fabric[x+leftRight][y+topBottom]++;
			}
		}
	}

	for _, s := range input {
		// split the given string into characters
		//use spaces to delimit since triming is weird with t prefix and suffix
		splitSpace := strings.Split(s, " ");
		// claimId := strings.Trim(splitSpace[0], "#")
		// leftRight at 0, topBottom at 1
		edge := strings.Split(splitSpace[2], ",")
		widthHeight := strings.Split(splitSpace[3], "x")
		//width at 0, height at 1
		leftRight, err := strconv.Atoi(edge[0])
		if err != nil {
			// handle error
		 }
		// fmt.Print(leftRight);
		topBottom, err := strconv.Atoi(strings.Replace(edge[1], ":", "", -1))
		if err != nil {
			// handle error
		 }
		width, err := strconv.Atoi(widthHeight[0])
		if err != nil {
			// handle error
		 }
		height, err := strconv.Atoi(widthHeight[1])
		if err != nil {
			// handle error
		 }
		//start at shift
		//find the fabric that never intersects
		//super sloppy
		for x := 0; x < width; x++ {
			for y:= 0; y < height; y++ {
				// mark all items in the square with the shift
				if fabric[x+leftRight][y+topBottom] > 1 {
					overlaps[intersect] = true
					break
				}
			}
			if overlaps[intersect] {
				break
			}
		}
		intersect++
	}

	for i:= 0; i < 1295; i++ {
		//find where there is no overlap
		if overlaps[i] == false {
			//the id is index+1 there should only be one
			fmt.Print(i+1);
			fmt.Print("\n")
			break
		}
	}
}