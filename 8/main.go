package main

import "fmt"

func main() {
	// int64 number
	var number int64
	number = 12345678912345
	// print number before changing
	fmt.Printf("Number: %d, number in binary Before changing: %b\n", number, number)
	i := 10
	// create mask for XOR
	mask := 1 << (i - 1)
	fmt.Printf("Number: %d, number in binary After changing: %b\n", number, number^int64(mask))
}
