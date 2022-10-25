package main

import "fmt"

func deleteElemByIndex(index *int, source_slice []int) (new_slice []int) {
	// make new slice with append func
	new_slice = append(source_slice[:*index], source_slice[*index+1:]...)
	return
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 6}
	index := 2
	fmt.Println(deleteElemByIndex(&index, slice))
}
