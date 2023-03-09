package main

import (
	"fmt"
	"sync"
	"time"
)

// we only send 3 values, but, main goroutine is waiting 4 values
// in that moment, only one goroutine is running: main
func main() {
	ch := make(chan int, 4)
	wg := &sync.WaitGroup{}
	wg.Add(3)

	go sendData(1, ch, wg)
	go sendData(2, ch, wg)
	go sendData(3, ch, wg)

	//if there are other goroutines running, the receiver waits
	//but when there aren't goroutines running, deadlock happen
	go onlyWaiting()

	fmt.Printf("Sleeping...\n")
	time.Sleep(time.Second * 2)
	fmt.Printf("Waking up...\n")

	showData(ch)
	showData(ch)
	showData(ch)
	showData(ch)

	wg.Wait()
}

func sendData(num int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- num
}

func showData(ch chan int) {
	fmt.Printf("Received: %v\n", <-ch)
}

func onlyWaiting() {
	time.Sleep(time.Second * 10)
}
