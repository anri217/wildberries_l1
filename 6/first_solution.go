package main

import "fmt"

func main() {
	// make chan bool
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			}
		}
	}()
	// send mes to chan and finish goroutine
	quit <- true
	fmt.Println("Finish goroutine")
	close(quit)
}
