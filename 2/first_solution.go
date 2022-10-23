package main

import (
	"fmt"
	"math"
	"sync"
)

func mathPow(num float64, wg *sync.WaitGroup) {
	// wg.Done() decrements the WG counter by one
	defer wg.Done()
	fmt.Println(math.Pow(num, 2))
}

func main() {
	numbers := []float64{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, num := range numbers {
		// raise WG counter
		wg.Add(1)
		go mathPow(num, &wg)
	}
	// waiting for all goroutines
	wg.Wait()
}
