package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	header             [2]int
	numChildNodes      int
	numMetaDataEntries int
}

func main() {
	// fileHandle, _ := os.Open("day8input.txt")
	fileHandle, _ := os.Open("day8testinput.txt")
	fileScanner := bufio.NewScanner(fileHandle)
	input := []int{}
	for fileScanner.Scan() {
		//split and remove spaces then populate input
		scan := strings.Split(fileScanner.Text(), " ")
		for i := 0; i < len(scan); i++ {
			b, _ := strconv.Atoi(scan[i])
			input = append(input, b)
		}
	}
	//remove spaces
	fmt.Println(input)
	firstHeader := [2]int{input[0], input[1]}
	childNodes := firstHeader[0]
	metadataEntries := firstHeader[1]
	//start at the second index
	startNode := node{firstHeader, childNodes, metadataEntries}
	sum, output := performRecursion(0, input, startNode, 2)
	//get the last sum since it wasnt originally included
	for i := (len(input) - metadataEntries); i < len(input); i++ {
		sum += input[i]
	}
	// firstTwo := [2]int(input[0]-'0'),input[1]-'0'}
	// firstTwo := [2]int(input[0]-'0'),input[1]-'0'}
	// startNode =
	fmt.Println("Output", output)
	fmt.Println("Sum of metadata", sum)
	fmt.Println("Test")
}

//return the current sum, the total input and the current node
func performRecursion(sum int, input []int, n node, currentIndex int) (int, node) {
	//get the child nodes
	//add the first child to the list
	tempNode := n
	tempNode.header = [2]int{input[currentIndex], input[currentIndex+1]}
	tempNode.numChildNodes = tempNode.header[0]
	tempNode.numChildNodes = tempNode.header[1]
	currentIndex += 2
	//check if the number of child nodes is 0
	//if so return back to the head and get the next header after summing the child
	if tempNode.header[0] == 0 {
		for i := 0; i < tempNode.header[1]; i++ {
			sum += input[currentIndex]
			fmt.Println(sum)
			currentIndex++
		}
		currentIndex++
		//recursively call back to the previous head with the new index and sum
		n.numChildNodes--
		n.numMetaDataEntries = 0
		return performRecursion(sum, input, n, currentIndex)
	}
	//since number of children is not 0 get the next input, add to any previous child nodes
	//get the number of childnodes
	//get the next child
	if tempNode.numChildNodes != 0 {
		//if the number of children is not 0 get the next child
		tempNode.numChildNodes = tempNode.header[0]
		//this will return the number of meta data entries
		tempNode.numMetaDataEntries = tempNode.header[1]
		//increase the index +1
		currentIndex++
		if tempNode.numChildNodes != 0 {
			return performRecursion(sum, input, tempNode, currentIndex)
		}
		return performRecursion(sum, input, n, currentIndex)
	}
	//get the meta data
	for i := 0; i < n.numMetaDataEntries; i++ {
		sum += input[currentIndex]
	}
	n.numMetaDataEntries = 0

	return sum, n
}
