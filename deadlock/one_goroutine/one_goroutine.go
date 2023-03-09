package main

import "fmt"

// cause ch is synchronous, it is waiting until sender and receiver are ready
// if ch is asynchronous, it will not block
func main() {
	ch := make(chan int)

	//send value to channel
	ch <- 10

	//receive value from channel
	g1 := <-ch

	fmt.Println("get g1: ", g1)
}
