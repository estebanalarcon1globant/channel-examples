package main

import (
	"fmt"
	"time"
)

// mientras el channel siga abierto, seguirá funcionando la operacion for
func main() {
	ch := make(chan int, 6)
	go sendIntegers(ch)

	for val := range ch {
		fmt.Println("Value: ", val)
	}

	fmt.Println("Out of range")
}

func sendIntegers(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	time.Sleep(4 * time.Second)
	close(ch)
	fmt.Println("Channel closed")
}
