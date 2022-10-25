package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	ch := make(chan int)

	// add wg counter
	wg.Add(1)

	go func() {
		for {
			foo, ok := <-ch
			if !ok {
				defer wg.Done()
				println("Finish goroutine")
				return
			}
			println(foo)
		}
	}()

	// write to channel
	for i := 0; i < 5; i++ {
		ch <- rand.Int()
	}

	close(ch)

	// wait for all goroutines
	wg.Wait()
}
