package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("day8input.txt")
	if err != nil {
		panic(err)
	}
	solve1(string(f))
}

func solve1(in string) {
	var ints []int

	for _, d := range strings.Split(strings.TrimSpace(in), " ") {
		i, _ := strconv.Atoi(d)
		ints = append(ints, i)
	}

	_, p1, p2 := parsenode(ints, 0)
	fmt.Println(p1, p2)
}

func parsenode(input []int, pos int) (newPos, sum, nodeval int) {
	childs := input[pos]
	metadatas := input[pos+1]
	pos += 2

	var childSums []int
	for c := 0; c < childs; c++ {
		var incsum int
		var chsum int
		pos, incsum, chsum = parsenode(input, pos)
		sum += incsum
		childSums = append(childSums, chsum)
	}

	refSums := 0
	sumOfMeta := 0

	for m := 0; m < metadatas; m++ {
		meta := input[pos]
		sum += meta
		sumOfMeta += meta
		if meta > 0 && meta < len(childSums)+1 {
			refSums += childSums[meta-1]
		}
		pos++
	}

	if childs == 0 {
		return pos, sum, sumOfMeta
	}

	return pos, sum, refSums
}
