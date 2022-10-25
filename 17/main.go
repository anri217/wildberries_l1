package main

import "fmt"

func binarySearch(needle int, haystack []int) bool {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		// init middle index in arr
		median := (low + high) / 2

		// if our search value is bigger than middle elem - we cut left side of arr and then continue searching in right part of arr while low <= high
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	// in this search we use only sorted arrays - thats why if low bigger than length - array do not include our val
	// and if the low index is not an index of our searching val (after binary search) - array do not include this val
	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))
}
