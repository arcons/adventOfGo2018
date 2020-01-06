package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//populate the treemap
	fileHandle, _ := os.Open("day25.txt")
	fileScanner := bufio.NewScanner(fileHandle)
	counter := 0
	starMap := make([][]int, 4)
	for i := range starMap {
		starMap[i] = make([]int, 1006)
	}
	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), ",")
		for i := 0; i < 4; i++ {
			t, _ := strconv.Atoi(s[i])
			starMap[counter][i] = t
		}
		counter++
	}

	fmt.Print("Just finish this one")
}
