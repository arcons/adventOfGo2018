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
	registers [3]int
}

//check for behaving like 3 opcodds
func main() {

	input := [37]string{}
	fileHandle, _ := os.Open("day19input.txt")
	fileScanner := bufio.NewScanner(fileHandle)
	counter := 0
	for fileScanner.Scan() {
		input[counter] = fileScanner.Text()
		counter++
	}

	inputSplit := strings.Split(input[0], " ")
	instructionPointer, _ := strconv.Atoi(inputSplit[1])
	updatedInput := remove(input, 0)
	outputReg := [6]int{}
	ip := 0
	for true {
		if ip > 35 || outputReg[instructionPointer] < 0 {
			break
		}
		outputReg, ip = performOperation(processInput(updatedInput[ip]), outputReg, instructionPointer)
		ip = outputReg[instructionPointer]
		ip++
		// ip = registers[init_ip]
		// while ip < len(code):
		//    registers[init_ip] = ip
		//    fnmap[code[ip][0]](registers, code[ip])
		//    ip = registers[init_ip]
		//    ip +=1
	}
	//instruction pointer at register 3

	fmt.Println("Part A: ", outputReg[0])
}

func remove(slice [37]string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func performOperation(op instruction, val [6]int, ip int) ([6]int, int) {
	outputReg := val
	regAVal := op.registers[0]
	regBVal := op.registers[1]
	regCVal := op.registers[2]
	if op.opcode == "addr" {
		outputReg[regCVal] = (val[regAVal] + val[regBVal])

	} else if op.opcode == "addi" {
		outputReg[regCVal] = (val[regAVal] + regBVal)

	} else if op.opcode == "mulr" {
		outputReg[regCVal] = (val[regAVal] * val[regBVal])

	} else if op.opcode == "muli" {
		outputReg[regCVal] = (val[regAVal] * regBVal)

	} else if op.opcode == "banr" {
		outputReg[regCVal] = (outputReg[regAVal] & outputReg[regBVal])

	} else if op.opcode == "bani" {
		outputReg[regCVal] = (outputReg[regAVal] & regBVal)

	} else if op.opcode == "borr" {
		outputReg[regCVal] = (outputReg[regAVal] | outputReg[regBVal])

	} else if op.opcode == "bori" {
		outputReg[regCVal] = (outputReg[regAVal] | regBVal)

	} else if op.opcode == "setr" {
		outputReg[regCVal] = outputReg[regAVal]

	} else if op.opcode == "seti" {
		outputReg[regCVal] = regAVal

	} else if op.opcode == "gtir" {
		if regAVal > outputReg[regBVal] {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if op.opcode == "gtri" {
		if outputReg[regAVal] > regBVal {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if op.opcode == "gtrr" {
		if outputReg[regAVal] > outputReg[regBVal] {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if op.opcode == "eqir" {
		if regAVal == outputReg[regBVal] {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if op.opcode == "eqri" {
		if outputReg[regAVal] == regBVal {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if op.opcode == "eqrr" {
		if outputReg[regAVal] == outputReg[regBVal] {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}
	}
	fmt.Println(op, outputReg, ip)
	return outputReg, ip
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
	registers := [3]int{reg1, reg2, reg3}
	output.registers = registers
	return output
}
