package main

import (
	"fmt"
	"math"
)

// input in channel
func mathPow(num *float64, ch *chan float64) {
	*ch <- math.Pow(*num, 2)
}

func main() {
	ch := make(chan float64)
	numbers := []float64{2, 4, 6, 8, 10}

	for _, num := range numbers {
		go mathPow(&num, &ch)

		// output from channel
		fmt.Println(<-ch)
	}
}
