package main

import (
	"container/ring"
	"fmt"
)

//internet ring example
func main() {
	const numPlayers = 439
	lastMarble := 7130700
	// const numPlayers = 17
	// lastMarble := 1104
	fmt.Println(numPlayers, lastMarble)
	scores := []int{}
	scores = make([]int, numPlayers)
	// create a ring and populate it with some values
	r := ring.New(1)
	r.Value = 0
	//populate the ring
	fmt.Println(r)
	for i := 1; i <= lastMarble; i++ {
		//check if its the 23rd marble
		if i%23 == 0 {
			//go back 8 spaces, 0 is in the middle
			//was reinitializing with colon here
			r = r.Move(-8)
			removeMe := r.Unlink(1)
			scores[i%numPlayers] += i + removeMe.Value.(int)
			r = r.Move(1)
		} else {
			//set r to its next space and add the value
			r = r.Move(1)
			tempRing := ring.New(1)
			tempRing.Value = i
			r.Link(tempRing)
			r = r.Move(1)
		}
		//use the move method to move back 7 spaces then pop off the marble 7 back
	}

	//output the high score
	highScore := 0
	for i := range scores {
		if scores[i] > highScore {
			highScore = scores[i]
		}
	}
	fmt.Println(highScore)
}
