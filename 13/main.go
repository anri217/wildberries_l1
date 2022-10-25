package main

import "fmt"

// too easy))
func main() {
	a, b := 5, 6
	fmt.Printf("a: %v, b: %v\n", a, b)

	a, b = b, a
	fmt.Printf("a: %v, b: %v\n", a, b)
}
