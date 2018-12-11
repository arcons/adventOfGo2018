package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type pair struct {
	letter byte
}

func fail() {
	fileHandle, _ := os.Open("day7testinput.txt")
	// input := [101]string{}
	var counter = 0
	fileScanner := bufio.NewScanner(fileHandle)
	listOfInput := list.New()
	// add first element
	scan := strings.Split(fileScanner.Text(), " ")
	beforeLetter := []byte(scan[1])
	afterLetter := []byte(scan[7])
	// push in the first two letters
	listOfInput.PushBack(&pair{beforeLetter[0]})
	listOfInput.PushBack(&pair{afterLetter[0]})
	//definitely a clean way of doing this
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		scan := strings.Split(fileScanner.Text(), " ")
		beforeLetter := []byte(scan[1])
		afterLetter := []byte(scan[7])
		// inputPair := &pair{beforeLetter[0]}
		// listOfInput.InsertAfter(pair{afterLetter}, *pair{beforeLetter})
		fmt.Print(beforeLetter[0])
		// fmt.Print("<-")
		fmt.Print(afterLetter[0])
		counter++
	}
	for e := listOfInput.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		fmt.Println(e.Value)
	}
}
