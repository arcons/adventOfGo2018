package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	opcode    int
	registers [4]int
}

//check for behaving like 3 opcodds
func main() {

	input := [4378]string{}
	fileHandle, _ := os.Open("day16input.txt")
	fileScanner := bufio.NewScanner(fileHandle)
	counter := 0
	beforeInstruction := [815]instruction{}
	// beforeInstruction = make([]instruction, 815)

	instructionOperation := [815]instruction{}
	// instructionOperation = make([]instruction, 815)

	afterInstruction := [815]instruction{}
	allInstructions := [815][3]instruction{}
	// afterInstruction = make([]instruction, 815)
	for fileScanner.Scan() {
		input[counter] = fileScanner.Text()
		counter++
	}
	counter = 0
	morethanthree := 0
	operations := []string{}
	for i := 0; i < 815; i++ {
		beforeInstruction[i] = processInput(input[counter])
		instructionOperation[i] = processInput(input[counter+1])
		afterInstruction[i] = processInput(input[counter+1])
		inputInstruction := [3]instruction{beforeInstruction[i], instructionOperation[i], afterInstruction[i]}
		allInstructions[i] = inputInstruction
		numOfPossible := 0
		operations = checkOperation(allInstructions[i])
		for j := range operations {
			if operations[j] != "" {
				numOfPossible++
			}
		}
		if numOfPossible > 2 {
			morethanthree++
		}
		counter += 4
		// fmt.Println(counter)
	}

	fmt.Println("Output: ", morethanthree)
}

func checkOperation(input [3]instruction) []string {
	var operations = []string{}
	add := checkAdd(input)
	mult := checkMulti(input)
	and := checkAnd(input)
	or := checkOr(input)
	set := checkSet(input)
	great := checkGreater(input)
	equal := checkEqual(input)
	operations = append(operations, add[0], add[1])
	operations = append(operations, mult[0], mult[1])
	operations = append(operations, and[0], and[1])
	operations = append(operations, or[0], or[1])
	operations = append(operations, set[0], set[1])
	operations = append(operations, great[0], great[1], great[2])
	operations = append(operations, equal[0], equal[1], equal[2])

	return operations
}

func checkAdd(input [3]instruction) [2]string {
	output := [2]string{}
	regA := input[0].registers[1]
	regB := input[0].registers[2]
	if (regA + regB) == input[1].registers[3] {
		output[0] = "addr"
	}
	if (input[1].registers[regA] + input[1].registers[2]) == input[1].registers[3] {
		output[1] = "addi"
	}
	return output
}

func checkMulti(input [3]instruction) [2]string {
	output := [2]string{}
	regA := input[0].registers[1]
	regB := input[1].registers[2]
	if (regA * regB) == input[1].registers[3] {
		output[0] = "mulr"
	}
	if (input[1].registers[regA] * input[1].registers[2]) == input[1].registers[3] {
		output[1] = "muli"
	}
	return output
}

func checkAnd(input [3]instruction) [2]string {
	output := [2]string{}
	regA := input[0].registers[1]
	regB := input[0].registers[2]
	if (input[1].registers[regA] & input[1].registers[regB]) == input[1].registers[3] {
		output[0] = "banr"
	}
	if (input[1].registers[regA] & input[1].registers[2]) == input[1].registers[3] {
		output[1] = "bani"
	}
	return output
}

func checkOr(input [3]instruction) [2]string {
	output := [2]string{}
	regA := input[0].registers[1]
	regB := input[0].registers[2]
	if (input[1].registers[regA] | input[0].registers[regB]) == input[1].registers[3] {
		output[0] = "borr"
	}
	if (input[1].registers[regA] | input[1].registers[2]) == input[1].registers[3] {
		output[1] = "bori"
	}
	return output
}

func checkSet(input [3]instruction) [2]string {
	output := [2]string{}
	regA := input[0].registers[1]
	if input[1].registers[regA] == input[1].registers[3] {
		output[0] = "setr"
	}
	if regA == input[1].registers[3] {
		output[1] = "seti"
	}
	return output
}

func checkGreater(input [3]instruction) [3]string {
	output := [3]string{}
	regA := input[0].registers[1]
	regB := input[0].registers[2]
	if input[1].registers[3] == 1 || input[1].registers[3] == 0 {
		if regA > input[0].registers[2] {
			output[0] = "gtir"
		}
		if input[0].registers[1] > input[1].registers[2] {
			output[1] = "gtri"
		}
		if regA > regB {
			output[2] = "gtrr"
		}
	}
	return output
}

func checkEqual(input [3]instruction) [3]string {
	output := [3]string{}
	//registers of the input
	regA := input[0].registers[1]
	regB := input[0].registers[2]
	if input[1].registers[3] == 1 || input[1].registers[3] == 0 {
		if input[1].registers[1] == regB {
			output[0] = "eqir"
		}
		if regA == input[1].registers[2] {
			output[1] = "eqri"
		}
		if regA == regB {
			output[2] = "eqrr"
		}
	}
	return output
}

func processInput(text string) instruction {
	inputSplit := strings.Split(text, " ")
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	output := instruction{}
	if strings.Contains(inputSplit[0], "Before:") || strings.Contains(inputSplit[0], "After") {
		output.opcode, _ = strconv.Atoi(reg.ReplaceAllString(inputSplit[1], ""))
		reg1, _ := strconv.Atoi(reg.ReplaceAllString(inputSplit[2], ""))
		reg2, _ := strconv.Atoi(reg.ReplaceAllString(inputSplit[3], ""))
		reg3, _ := strconv.Atoi(reg.ReplaceAllString(inputSplit[4], ""))
		registers := [4]int{0, reg1, reg2, reg3}
		output.registers = registers
	} else {
		output.opcode, _ = strconv.Atoi(inputSplit[0])
		reg1, _ := strconv.Atoi(inputSplit[1])
		reg2, _ := strconv.Atoi(inputSplit[2])
		reg3, _ := strconv.Atoi(inputSplit[3])
		registers := [4]int{0, reg1, reg2, reg3}
		output.registers = registers
	}
	return output
}
