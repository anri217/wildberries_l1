package main

import (
	"fmt"
	"math"
)

// write result to channel
func mathPow(ch *chan float64, num *float64) {
	*ch <- math.Pow(*num, 2)
}

func main() {
	ch := make(chan float64)
	numbers := []float64{2, 4, 6, 8, 10}
	var sum float64

	for _, num := range numbers {
		go mathPow(&ch, &num)
		// add square to the sum from the channel
		sum += <-ch
	}
	fmt.Println(sum)
}
