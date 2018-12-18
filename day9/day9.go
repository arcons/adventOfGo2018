package main

import(
	"fmt"
	"container/ring"
)

func main(){
	package main

import (
	"container/ring"
	"fmt"
	"os"
	"time"
)

//internet ring example
func main() {
	coffee := []string{"kenya", "guatemala", "ethiopia"}

	// create a ring and populate it with some values
	r := ring.New(len(coffee))
	for i := 0; i < r.Len(); i++ {
		r.Value = coffee[i]
		r = r.Next()
	}

	// print all values of the ring, easy done with ring.Do()
	r.Do(func(x interface{}) {
		fmt.Println(x)
	})

	// end after 10s otherwise play.golang.org gives an error
	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("-- exiting --")
		os.Exit(0)
	}()

	// .. or each one by one.
	for _ = range time.Tick(time.Second * 1) {
		r = r.Next()
		fmt.Println(r.Value)
	}
}
}