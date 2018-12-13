package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type rule struct {
	substring string
	output    byte
}

func main() {
	initialState := "#....##.#.#.####..#.######..##.#.########..#...##...##...##.#.#...######.###....#...##..#.#....##.##"
	const ruleNum = 32
	// const ruleNum = 14
	// initialState := "#..#.#..##......###...###"
	var offSet = 0
	var buffer bytes.Buffer
	buffer.WriteString(".....")
	buffer.WriteString(initialState)
	buffer.WriteString(".....")
	offSet += 5
	initialState = buffer.String()
	buffer.Reset()
	// number of rules
	fmt.Println(initialState)
	fileHandle, _ := os.Open("day12input.txt")
	fileScanner := bufio.NewScanner(fileHandle)
	counter := 0
	rules := [ruleNum]rule{}
	// read in all of the rules
	for fileScanner.Scan() {
		scan := strings.Split(fileScanner.Text(), " ")
		rules[counter] = rule{scan[0], []byte(scan[2])[0]}
		counter++
	}
	//(No other pots currently contain plants.) so everything behind and past the length of the initial state contain.
	// Make a byte array of size 500 and start at 250 then use 250 as 0 so I dont have to figure out how to dynamically aloate size
	// keep track of the position to update and their position in a map
	// 30697 2000 10005
	// 30712 2001
	// 30727 2002 10015
	// 2444
	tempState := []byte(initialState)
	outputState := make([]byte, len(tempState))
	prevSum := 0
	stringShiftIndex := 0
	for i := 0; i < 200; i++ {
		//check each index for eah rule
		// for j := 0; j < len(initialState); j++ {
		//check if the rule applies to the string
		//initalize all values
		outputState = make([]byte, len(tempState))
		copy(outputState, tempState)
		for k := 0; k < len(tempState); k++ {
			outputState[k] = '.'
		}
		for j := 0; j < len(tempState)-5; j++ {
			// if j == 26 && i == 4 {
			// 	fmt.Print("Stop here and step through")
			// }
			currentSection := string(tempState[j : 5+j])
			for _, r := range rules {
				if currentSection == r.substring {
					//the index needed to change is the middle index
					substringIndex := 2 + j
					//check if the item is too large
					outputState[substringIndex] = r.output
					if substringIndex+5 > len(tempState) {
						buffer.WriteString(".....")
						buffer.WriteString(string(outputState))
						buffer.WriteString(".....")
						offSet += 5
						outputState = []byte(buffer.String())
						// fmt.Println("increased buffer size")
						buffer.Reset()
					}
					// fmt.Printf("Contains substring %s at index %d\n", r.substring, substringIndex-offSet)
				}
			}
		}
		// println(string(outputState), i+1)

		sum := 0
		for n := 0; n < len(outputState); n++ {
			if outputState[n] == 35 {
				sum += (n - offSet)
			}
		}
		if strings.Contains(string(outputState), string(tempState)) {
			// if strings.Contains(string(tempState), string(outputState)) {
			fmt.Println(" ", i)
			stringShiftIndex = i
			fmt.Println(prevSum)
			fmt.Println(sum)
			fmt.Println(stringShiftIndex)
			break
		}
		prevSum = sum
		tempState = make([]byte, len(outputState))
		copy(tempState, outputState)
	}
	//at iteration 43, it starts to shift

	// fmt.Println(string(outputState))
	// fmt.Print(offSet)
	sum := 0
	// fmt.Println(string(outputState[offSet:(len(outputState) - offSet)]))
	for i := 0; i < len(outputState); i++ {
		if outputState[i] == 35 {
			sum += (i - offSet)
		}
	}
	shiftDiff := sum - prevSum
	lastIndex := 50000000000 - (int(stringShiftIndex))
	value := (lastIndex * shiftDiff) + prevSum
	fmt.Println(value)
	// 1005 3697
	// 1010 3712
	// incorrect

	//sum up all of the indexes with plans in them for the solution
	//subtract all number less than the offset then add those above the offset
}
