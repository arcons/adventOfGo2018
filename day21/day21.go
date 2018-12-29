package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	opcode    string
	registers [4]int
}

//check for behaving like 3 opcodds
func main() {

	input := [32]string{}
	fileHandle, _ := os.Open("day21input.txt")
	fileScanner := bufio.NewScanner(fileHandle)
	operations := [32]instruction{}
	counter := 0
	for fileScanner.Scan() {
		input[counter] = fileScanner.Text()
		counter++
	}

	outputReg := [4]int{0, 0, 0, 0}
	for i := 1; i < 32; i++ {
		operations[i] = processInput(input[i])
		outputReg = performOperation(operations[i], outputReg)
	}
	
	fmt.Println("Part 2: ", outputReg[0])
}

func performOperation(op instruction, val [4]int) [4]int {
	outputReg := val
	regAVal := op.registers[1]
	regBVal := op.registers[2]
	regCVal := 0
	if op.opcode == "addr" {
		outputReg[regCVal] = (val[regAVal] + val[regBVal])

	} else if op.opcode == "addi" {
		outputReg[regCVal] = (val[regAVal] + regBVal)

	} else if op.opcode == "mulr" {
		outputReg[regCVal] = (val[regAVal] * val[regBVal])

	} else if op.opcode == "muli" {
		outputReg[regCVal] = (val[regAVal] * regBVal)

	} else if op.opcode == "banr" {
		outputReg[regCVal] = (val[regAVal] & val[regBVal])

	} else if op.opcode == "bani" {
		outputReg[regCVal] = (val[regAVal] & regBVal)

	} else if op.opcode == "borr" {
		outputReg[regCVal] = (val[regAVal] | val[regBVal])

	} else if op.opcode == "bori" {
		outputReg[regCVal] = (val[regAVal] | regBVal)

	} else if op.opcode == "setr" {
		outputReg[regCVal] = val[regAVal]

	} else if op.opcode == "seti" {
		outputReg[regCVal] = regAVal

	} else if op.opcode == "gtir" {
		if regAVal > val[regBVal] {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if op.opcode == "gtri" {
		if val[regAVal] > regBVal {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if op.opcode == "gtrr" {
		if val[regAVal] > val[regBVal] {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if op.opcode == "eqir" {
		if regAVal == val[regBVal] {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if op.opcode == "eqri" {
		if val[regAVal] == regBVal {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if op.opcode == "eqrr" {
		if val[regAVal] == val[regBVal] {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}
	}

	return outputReg
}

func processInput(text string) instruction {
	inputSplit := strings.Split(text, " ")
	output := instruction{}

	// reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	output.opcode = inputSplit[0]
	reg1, _ := strconv.Atoi(inputSplit[1])
	reg2, _ := strconv.Atoi(inputSplit[2])
	reg3, _ := strconv.Atoi(inputSplit[3])
	registers := [4]int{0, reg1, reg2, reg3}
	output.registers = registers
	return output
}
