package main

import "fmt"

func intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	// for each a item we put true value in map
	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		// if we have such b item as a key from map - add to c as an intersection result
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func main() {
	a := []int{7, 2, 3, 4, 5}
	b := []int{2, 3, 6, 8}
	c := intersection(a, b)
	fmt.Println(c)
	// for string type we will do the same way
}
