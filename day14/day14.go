package main

import (
	"fmt"
)

func main() {
	// startValue := 890691
	// startValue := int64(37)
	// fmt.Println(startValue)
	//create an array of digits
	//start at first two digits from the start
	posA := 0
	posAValue := 0
	posB := 1
	posBValue := 0
	sum := 0
	arrayOfDigits := []int{}
	tempArrayOfDigits := []int{8, 9, 0, 6, 9, 1}

	// tempArrayOfDigits := []int{3, 7}
	// startDigit := 2
	for len(tempArrayOfDigits) < 20 {
		//get the values at the two positions
		numDigits := len(tempArrayOfDigits)
		//do this outside so memory allocation doesnt need to happen twice
		arrayOfDigits = make([]int, numDigits)
		// tempStart = startValue
		for j := 0; j < numDigits; j++ {
			arrayOfDigits[j] = tempArrayOfDigits[j]
		}
		//calculate the number of digits
		posAValue = arrayOfDigits[posA]
		posBValue = arrayOfDigits[posB]
		// fmt.Println(posAValue, posBValue)
		sum = arrayOfDigits[posA] + arrayOfDigits[posB]
		tempSum := sum
		sumDigits := 0
		for tempSum != 0 {
			tempSum /= 10
			sumDigits++
		}
		tempArrayOfDigits = arrayOfDigits
		numDigits += sumDigits
		//reallocate the memory for the array of digits
		arrayOfDigits = make([]int, numDigits)
		for j := 0; j < numDigits-sumDigits; j++ {
			//store the values again
			arrayOfDigits[j] = tempArrayOfDigits[j]
		}
		tempSum = sum
		for j := sumDigits - 1; j >= 0; j-- {
			currentInt := tempSum % 10
			tempSum /= 10
			arrayOfDigits[j+numDigits-sumDigits] = currentInt
		}
		posA = posA + posAValue + 1
		posB = posB + posBValue + 1
		if posA >= numDigits {
			posA = posA % (numDigits)
		}
		if posB >= numDigits {
			posB = posB % (numDigits)
		}
		tempArrayOfDigits = arrayOfDigits
	}
	// fmt.Println(posA, posB)
	fmt.Println(arrayOfDigits)
	start := 6
	for i := 0; i < 10; i++ {
		if start == len(arrayOfDigits) {
			start = 0
		}
		fmt.Print(arrayOfDigits[start])
		start++
	}
}
