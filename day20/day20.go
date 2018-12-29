package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileHandle, _ := os.Open("day20testinput.txt")
	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Scan()
	input := []byte(fileScanner.Text())
	fmt.Println("Day 20", input)
}
