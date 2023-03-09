package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	name  string
	value int
}

func main() {
	unbufferedChan := make(chan Task, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	//create task
	task := Task{
		name:  "testing task",
		value: 5,
	}

	//send data asynchronously
	go SendTask(task, unbufferedChan, wg)

	time.Sleep(time.Second * 4)
	fmt.Printf("I want to modify the task...\n")
	task.value = 10
	fmt.Printf("task modified...\n")

	//Receiving from channel
	fmt.Printf("Received from channel: %v\n", <-unbufferedChan)

	wg.Wait()
}

func SendTask(task Task, unbufferedChan chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Sending to channels: %v\n", task)
	unbufferedChan <- task
	fmt.Printf("Sent to channels: %v\n", task)
}
