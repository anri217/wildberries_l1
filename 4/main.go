package main

import (
	"fmt"
	"time"
)

func workerStart() {
	n := 9

	channel := make(chan int)

	// create n workers
	for i := 0; i < n; i++ {
		go worker(&i, &channel)
	}

	// write to the channel
	for {
		channel <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

func worker(worker *int, channel *chan int) {
	for {
		number := <-*channel
		fmt.Printf("goroutine #%d: value: %d\n", *worker, number)
	}
}

func main() {
	workerStart()
}
