package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	numbers := []float64{2, 4, 6, 8, 10}
	var sum float64
	var wg sync.WaitGroup

	for _, num := range numbers {
		// raise WG counter
		wg.Add(1)

		go func(num float64) {
			// wg.Done() decrements the WG counter by one
			defer wg.Done()
			sum += math.Pow(num, 2)
		}(num)
	}
	// wait for all goroutines
	wg.Wait()
	fmt.Println(sum)
}
