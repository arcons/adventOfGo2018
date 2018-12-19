package main

import (
	"container/ring"
	"fmt"
)

//internet ring example
func main() {
	// numPlayers := 439
	// lastMarble := 71307
	numPlayers := 9
	lastMarble := 25
	fmt.Println(numPlayers, lastMarble)

	// create a ring and populate it with some values
	r := ring.New(lastMarble)
	r.Value = 0
	//populate the ring
	fmt.Println(r)
	for i := 1; i <= lastMarble; i++ {
		r = r.Next()
		r.Value = i
		fmt.Println(r)
		//use the move method to move back 7 spaces then pop off the marble 7 back
	}
}
