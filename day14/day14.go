package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// startValue := 890691
	posA := 0
	posAValue := 0
	posB := 1
	posBValue := 0
	arrayOfDigits := []byte{'3', '7'}
	//use a buffer and a byte array
	// tempArrayOfDigits := []byte{'3', '7'}
	// startDigit := 2
	for len(arrayOfDigits) < 100000000 {
		//get the values at the two positions
		// numDigits := len(tempArrayOfDigits)
		// //do this outside so memory allocation doesnt need to happen twice
		// arrayOfDigits = make([]byte, numDigits)
		// // tempStart = startValue
		// for j := 0; j < numDigits; j++ {
		// 	arrayOfDigits[j] = tempArrayOfDigits[j]
		// }
		//calculate the number of digits
		//use a byte array for a proper append
		posAValue = int(arrayOfDigits[posA] - '0')
		posBValue = int(arrayOfDigits[posB] - '0')
		// fmt.Println(posAValue, posBValue)
		sum := []byte(strconv.Itoa(int(posAValue + posBValue)))
		// tempSum := sum
		// sumDigits := 0
		// for tempSum != 0 {
		// 	tempSum /= 10
		// 	sumDigits++
		// }
		arrayOfDigits = append(arrayOfDigits, sum...)
		// tempArrayOfDigits = arrayOfDigits
		// numDigits += sumDigits
		//reallocate the memory for the array of digits
		// arrayOfDigits = make([]byte, numDigits)
		// for j := 0; j < numDigits-sumDigits; j++ {
		// 	//store the values again
		// 	arrayOfDigits[j] = tempArrayOfDigits[j]
		// }
		// tempSum = sum
		// for j := sumDigits - 1; j >= 0; j-- {
		// 	currentInt := tempSum % 10
		// 	tempSum /= 10
		// 	arrayOfDigits[j+numDigits-sumDigits] = byte(currentInt) + '0'
		// }
		//posA & B are value of them plus current position + 1
		numDigits := len(arrayOfDigits)
		posA = posA + posAValue + 1
		posB = posB + posBValue + 1
		if posA >= numDigits {
			posA = posA % (numDigits)
		}
		if posB >= numDigits {
			posB = posB % (numDigits)
		}
		// tempArrayOfDigits = arrayOfDigits
	}
	// fmt.Println(posA, posB)
	// fmt.Println(arrayOfDigits)
	start := 890691
	for i := 0; i < 10; i++ {
		if start == len(arrayOfDigits) {
			start = 0
		}
		fmt.Print(string(arrayOfDigits[start]))
		start++
	}
	fmt.Println(" ")
	// find the subset for part b
	startstring := strconv.Itoa(890691)
	fmt.Println(strings.Index(string(arrayOfDigits), startstring))
}
