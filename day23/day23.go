package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type vec struct {
	x, y, z int
}

type bot struct {
	v vec
	r int
}

func main() {
	//populate the treemap
	fileHandle, _ := os.Open("day23.txt")
	fileScanner := bufio.NewScanner(fileHandle)
	counter := 0

	botArray := [1006]bot{}
	maxBot := bot{}
	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), " ")
		s2 := strings.Split(s[0], "<")
		positions := strings.Split(s2[1], ",")
		x, _ := strconv.Atoi(positions[0])
		y, _ := strconv.Atoi(positions[1])
		z, _ := strconv.Atoi(positions[2])
		curVec := vec{x, y, z}
		s3 := strings.Split(s[2], "=")
		posRange, _ := strconv.Atoi(s3[1])
		curBot := bot{curVec, posRange}
		botArray[counter] = curBot
		if curBot.r > maxBot.r {
			maxBot = curBot
		}
		fmt.Println(posRange, curVec)
		counter++
	}
	sum := 0
	for i := 0; i < 1006; i++ {
		if distance(maxBot.v, botArray[i].v) <= maxBot.r {
			sum++
		}
	}
	// too high
	fmt.Print(sum)
}

func distance(v1 vec, v2 vec) int {
	return int(math.Abs(float64(v2.x-v1.x)) + math.Abs(float64(v2.y-v1.y)) + math.Abs(float64(v2.z-v1.z)))
}
