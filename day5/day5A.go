package main

import (
	"fmt"
	// "math"
	"strings"
	// "sort"
	// "time"
	"bytes"
)

func day5A() {
	//do a file input
	input := "dabAcCaCBAcCcaDA"
	fmt.Print("Day 5 AOCC\n")

	//create guard map with their number and 60 int array
	//+1 every time they fall asleep at a given minute
	//keep track of largest +1 and guard number
	var inputLength = len(input)
	var buffer bytes.Buffer
	for i := 1; i < inputLength; i++ {
		char0 := int(input[i-1])
		char1 := int(input[i])
		// fmt.Print(string(input[i-1]))
		if char0-char1 == 32 {
			//remove the list and subtract 1 from i
			buffer.WriteString(string(input[i-1]))
			buffer.WriteString(string(input[i]))
			removeString := buffer.String()
			// fmt.Print(removeString)
			// fmt.Print("\n")
			input = strings.Replace(input, removeString, "", 1)
			i -= 2
		} else if char1-char0 == 32 {
			//remove the list and subtract 1 from i
			buffer.WriteString(string(input[i-1]))
			buffer.WriteString(string(input[i]))
			removeString := buffer.String()
			// fmt.Print(removeString)
			// fmt.Print("\n")
			input = strings.Replace(input, removeString, "", 1)
			i -= 2
		}
		inputLength = len(input)
		buffer.Reset()
	}

	fmt.Print(input)
}
