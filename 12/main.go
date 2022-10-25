package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "three"}

	set := map[string]bool{}

	for _, cur_str := range arr {
		set[cur_str] = true
	}

	// only uniaue values will be here
	fmt.Println(set)
}
