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
	//operations for a given opcode
	operations := [815][]string{}
	opcodes := [16][]string{}
	for i := 0; i < 815; i++ {
		beforeInstruction[i] = processInput(input[counter])
		instructionOperation[i] = processInput(input[counter+1])
		afterInstruction[i] = processInput(input[counter+2])
		inputInstruction := [3]instruction{beforeInstruction[i], instructionOperation[i], afterInstruction[i]}
		allInstructions[i] = inputInstruction
		numOfPossible := 0
		operations[i] = checkOperation(allInstructions[i])
		for j := range operations[i] {
			if operations[i][j] != "" {
				numOfPossible++
			}
			opcodes[instructionOperation[i].opcode] = operations[i]
		}
		if numOfPossible > 2 {
			morethanthree++
		}

		//opcodes
		//11 = eqri
		counter += 4
		// fmt.Println(counter)
	}

	fmt.Println("Part 1: ", morethanthree)
	//delete the empty itmes
	for i := range opcodes {
		opcodes[i] = deleteEmpty(opcodes[i])
	}

	//build the map by looking at it,
	//10 eqrr
	//11 eqri
	//7  eqir
	//13 gtrr
	//15 gtri
	//5  setr
	//4  gtir
	//2  banr
	//3  bani
	//8  seti
	//14 mulr
	//1  muli
	//0  bori
	//12 borr
	//9  addi
	//6  addr
	// opMap := [16]string{"bori", "muli", "banr", "bani", "gtir", "setr", "addr", "eqir", "seti", "addi", "eqrr", "eqri", "borr", "gtrr", "mulr", "gtri"}

	// for i := range opMap {
	// 	fmt.Println(opMap[i])
	// }

	outputReg := [4]int{0, 0, 0, 0}
	for i := 3262; i < len(input); i++ {
		register := processInput(input[i])
		outputReg = performOperation(register, outputReg)
		//perform the operation and store it into outputReg
	}
	// register := processInput(input[4377])

	fmt.Println("Part 2: ", outputReg[0])
}

func performOperation(op instruction, val [4]int) [4]int {
	opMap := [16]string{"bori", "muli", "banr", "bani", "gtir", "setr", "addr", "eqir", "seti", "addi", "eqrr", "eqri", "borr", "gtrr", "mulr", "gtri"}
	outputReg := val
	regAVal := op.registers[1]
	regBVal := op.registers[2]
	regCVal := op.registers[3]
	if opMap[op.opcode] == "addr" {
		outputReg[regCVal] = (val[regAVal] + val[regBVal])

	} else if opMap[op.opcode] == "addi" {
		outputReg[regCVal] = (val[regAVal] + regBVal)

	} else if opMap[op.opcode] == "mulr" {
		outputReg[regCVal] = (val[regAVal] * val[regBVal])

	} else if opMap[op.opcode] == "muli" {
		outputReg[regCVal] = (val[regAVal] * regBVal)

	} else if opMap[op.opcode] == "banr" {
		outputReg[regCVal] = (val[regAVal] & val[regBVal])

	} else if opMap[op.opcode] == "bani" {
		outputReg[regCVal] = (val[regAVal] & regBVal)

	} else if opMap[op.opcode] == "borr" {
		outputReg[regCVal] = (val[regAVal] | val[regBVal])

	} else if opMap[op.opcode] == "bori" {
		outputReg[regCVal] = (val[regAVal] | regBVal)

	} else if opMap[op.opcode] == "setr" {
		outputReg[regCVal] = val[regAVal]

	} else if opMap[op.opcode] == "seti" {
		outputReg[regCVal] = regAVal

	} else if opMap[op.opcode] == "gtir" {
		if regAVal > val[regBVal] {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if opMap[op.opcode] == "gtri" {
		if val[regAVal] > regBVal {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if opMap[op.opcode] == "gtrr" {
		if val[regAVal] > val[regBVal] {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if opMap[op.opcode] == "eqir" {
		if regAVal == val[regBVal] {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if opMap[op.opcode] == "eqri" {
		if val[regAVal] == regBVal {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}

	} else if opMap[op.opcode] == "eqrr" {
		if val[regAVal] == val[regBVal] {
			outputReg[regCVal] = 1
		} else {
			outputReg[regCVal] = 0
		}
	}

	return outputReg
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		// this can be done by looping and removing but I'm lazy
		if str != "" && str != "mulr" && str != "muli" && str != "bori" && str != "banr" && str != "seti" && str != "bani" && str != "eqrr" && str != "eqri" && str != "eqir" && str != "gtrr" && str != "gtri" && str != "setr" && str != "gtir" {
			// if str != "" {
			r = append(r, str)
		}
	}
	return r
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
	// reg1Val := input[0].registers[1]
	// reg2Val := input[0].registers[2]
	regAVal := input[1].registers[1]
	regBVal := input[1].registers[2]
	regCVal := input[1].registers[3]
	// if (reg1Val + reg2Val) == regCVal {
	// 	output[0] = "addr"
	// }
	// if (reg1Val + input[0].registers[regBVal]) == regCVal {
	// 	output[1] = "addi"
	// }
	if (input[0].registers[regAVal] + input[0].registers[regBVal]) == input[2].registers[regCVal] {
		output[0] = "addr"
	}
	if (input[0].registers[regAVal] + regBVal) == input[2].registers[regCVal] {
		output[1] = "addi"
	}
	return output
}

func checkMulti(input [3]instruction) [2]string {
	output := [2]string{}
	// reg1Val := input[0].registers[1]
	// reg2Val := input[0].registers[2]
	regAVal := input[1].registers[1]
	regBVal := input[1].registers[2]
	regCVal := input[1].registers[3]
	if (input[0].registers[regAVal] * input[0].registers[regBVal]) == input[2].registers[regCVal] {
		output[0] = "mulr"
	}
	if (input[0].registers[regAVal] * regBVal) == input[2].registers[regCVal] {
		output[1] = "muli"
	}
	return output
}

func checkAnd(input [3]instruction) [2]string {
	output := [2]string{}
	// reg2Val := input[0].registers[2]
	// regA := input[1].registers[1]
	// regB := input[1].registers[2]
	// outputReg := input[1].registers[3]
	regAVal := input[1].registers[1]
	regBVal := input[1].registers[2]
	regCVal := input[1].registers[3]
	if (input[0].registers[regAVal] & input[0].registers[regBVal]) == input[2].registers[regCVal] {
		output[0] = "banr"
	}
	if (input[0].registers[regAVal] & regBVal) == input[2].registers[regCVal] {
		output[1] = "bani"
	}
	return output
}

func checkOr(input [3]instruction) [2]string {
	output := [2]string{}
	// reg2Val := input[0].registers[2]
	// regA := input[1].registers[1]
	// regB := input[1].registers[2]
	// outputReg := input[1].registers[3]
	// if (regA | regB) == outputReg {
	// 	output[0] = "borr"
	// }
	// if (regA | reg2Val) == outputReg {
	// 	output[1] = "bori"
	// }
	regAVal := input[1].registers[1]
	regBVal := input[1].registers[2]
	regCVal := input[1].registers[3]
	if (input[0].registers[regAVal] | input[0].registers[regBVal]) == input[2].registers[regCVal] {
		output[0] = "borr"
	}
	if (input[0].registers[regAVal] | regBVal) == input[2].registers[regCVal] {
		output[1] = "bori"
	}
	return output
}

func checkSet(input [3]instruction) [2]string {
	output := [2]string{}
	regAVal := input[1].registers[1]
	regCVal := input[1].registers[3]
	if input[0].registers[regAVal] == input[2].registers[regCVal] {
		output[0] = "setr"
	}
	if regAVal == input[2].registers[regCVal] {
		output[1] = "seti"
	}
	return output
}

func checkGreater(input [3]instruction) [3]string {
	output := [3]string{}
	// reg1Val := input[0].registers[1]
	// reg2Val := input[0].registers[2]
	regAVal := input[1].registers[1]
	regBVal := input[1].registers[2]
	regCVal := input[1].registers[3]
	if input[2].registers[regCVal] == 1 {
		if regAVal > input[0].registers[regBVal] {
			output[0] = "gtir"
		}
		if input[0].registers[regAVal] > regBVal {
			output[1] = "gtri"
		}
		if input[0].registers[regAVal] > input[0].registers[regBVal] {
			output[2] = "gtrr"
		}
	}
	if input[2].registers[regCVal] == 0 {
		if regAVal <= input[0].registers[regBVal] {
			output[0] = "gtir"
		}
		if input[0].registers[regAVal] <= regBVal {
			output[1] = "gtri"
		}
		if input[0].registers[regAVal] <= input[0].registers[regBVal] {
			output[2] = "gtrr"
		}

	}
	return output
}

func checkEqual(input [3]instruction) [3]string {
	output := [3]string{}
	//registers of the input
	// reg1Val := input[0].registers[1]
	// reg2Val := input[0].registers[2]
	regAVal := input[1].registers[1]
	regBVal := input[1].registers[2]
	regCVal := input[1].registers[3]
	if input[2].registers[regCVal] == 1 {
		if regAVal == input[0].registers[regBVal] {
			output[0] = "eqir"
		}
		if input[0].registers[regAVal] == regBVal {
			output[1] = "eqri"
		}
		if input[0].registers[regAVal] == input[0].registers[regBVal] {
			output[2] = "eqrr"
		}
	}
	if input[2].registers[regCVal] == 0 {
		if regAVal != input[0].registers[regBVal] {
			output[0] = "eqir"
		}
		if input[0].registers[regAVal] != regBVal {
			output[1] = "eqri"
		}
		if input[0].registers[regAVal] != input[0].registers[regBVal] {
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
	if strings.Contains(inputSplit[0], "Before:") {
		output.opcode, _ = strconv.Atoi(reg.ReplaceAllString(inputSplit[1], ""))
		reg1, _ := strconv.Atoi(reg.ReplaceAllString(inputSplit[2], ""))
		reg2, _ := strconv.Atoi(reg.ReplaceAllString(inputSplit[3], ""))
		reg3, _ := strconv.Atoi(reg.ReplaceAllString(inputSplit[4], ""))
		registers := [4]int{output.opcode, reg1, reg2, reg3}
		output.registers = registers
	} else if strings.Contains(inputSplit[0], "After") {
		output.opcode, _ = strconv.Atoi(reg.ReplaceAllString(inputSplit[2], ""))
		reg1, _ := strconv.Atoi(reg.ReplaceAllString(inputSplit[3], ""))
		reg2, _ := strconv.Atoi(reg.ReplaceAllString(inputSplit[4], ""))
		reg3, _ := strconv.Atoi(reg.ReplaceAllString(inputSplit[5], ""))
		registers := [4]int{output.opcode, reg1, reg2, reg3}
		output.registers = registers
	} else {
		output.opcode, _ = strconv.Atoi(inputSplit[0])
		reg1, _ := strconv.Atoi(inputSplit[1])
		reg2, _ := strconv.Atoi(inputSplit[2])
		reg3, _ := strconv.Atoi(inputSplit[3])
		registers := [4]int{output.opcode, reg1, reg2, reg3}
		output.registers = registers
	}
	return output
}
