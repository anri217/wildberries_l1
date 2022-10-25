package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(i int, ctx context.Context, channel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case num := <-channel:
			fmt.Printf("Worker № %d random Int value: %d\n", i, num)
		case <-ctx.Done():
			fmt.Printf("Close Worker stream № %d\n", i)
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 9

	channel := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	os_channel := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(os_channel, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup

	// create n workers
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i, ctx, channel, &wg)
	}

	// write to the channel
	for {
		select {
		// if we have mes for interrupt from OS - we do it
		case <-os_channel:
			fmt.Println("\nGraceful shutdown")
			cancel()
			wg.Wait()
			return
		// when we do not have mes - we wait second and send info
		case <-time.After(time.Second):
			channel <- rand.Int()
		}
	}
}
