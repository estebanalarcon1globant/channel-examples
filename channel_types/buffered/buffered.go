package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	unbufferedChan := make(chan string, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	//send data asynchronously
	go sendDataToBufChan("testing channel", unbufferedChan, wg)

	fmt.Printf("Sleeping...\n")
	time.Sleep(time.Second * 4)
	fmt.Printf("Waking up...\n")

	//Receiving from channel
	fmt.Printf("Received from channel: %v\n", <-unbufferedChan)

	wg.Wait()
}

func sendDataToBufChan(s string, unbufferedChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Sending to channels: %v\n", s)
	unbufferedChan <- s
	fmt.Printf("Sent to channels: %v\n", s)
}
