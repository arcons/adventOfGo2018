package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Test comment
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input ints: ")
	text, _ := reader.ReadString('\n')
	// perform the parsing
	if text == "1a" {
		day1A()
	}
	if text == "1b" {
		day1B()
	} else {
		fmt.Print("enter a or b to run the code")
	}

}
