package main

import "fmt"

// write nums to chan
func writer(array []int, ch1 *chan int) {
	for _, val := range array {
		*ch1 <- val
	}
	close(*ch1)
}

// read nums from chan and multiply
func reader(ch2 *chan int, ch1 *chan int) {
	for val := range *ch1 {
		*ch2 <- val * 2
	}
	close(*ch2)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	array := []int{1, 2, 3, 4, 5}

	go writer(array, &ch1)
	go reader(&ch2, &ch1)

	// вывод числе из канала 2 (val(ch1) * 2)
	for val := range ch2 {
		fmt.Println("x * 2: ", val)
	}
}
