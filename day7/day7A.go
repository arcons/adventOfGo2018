package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type letterPair struct {
	beforeLetter byte
	afterLetter  []byte
}
type Pairs map[byte]*letterPair
//screw this im doing python
func main() {
	fileHandle, _ := os.Open("day7testinput.txt")
	// input := [101]string{}
	var counter = 0
	fileScanner := bufio.NewScanner(fileHandle)
	m := make(Pairs)
	//definitely a clean way of doing this
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		scan := strings.Split(fileScanner.Text(), " ")
		beforeLetter := []byte(scan[1])
		afterLetter := []byte(scan[7])
		//check if the letter is in the list
		if len(m[beforeLetter[0]].afterLetter) == 0 {
			inputPair := letterPair{beforeLetter[0], afterLetter}
			m[beforeLetter[0]] = &inputPair
		} else {
			s := append(m[beforeLetter[0]].afterLetter, beforeLetter[0])
			m[beforeLetter[0]].afterLetter = s
		}
		// m[beforeLetter[0]] = inputPair
		// fmt.Println(inputPair)
		// fmt.Print(afterLetter[0])
		counter++
	}
	for _, k := range m {
		fmt.Println("Key:", k.beforeLetter)
	}
}
