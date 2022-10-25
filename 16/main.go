package main

import (
	"fmt"
	"math/rand"
	"time"
)

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Do it for two parts of current arr
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}

func main() {
	rand.Seed(time.Now().Unix())
	not_sorted := []int{5, 6, 1, 12, 6, 2435, 234, 120}
	fmt.Println("Sorted array: ", qsort(not_sorted))
}
