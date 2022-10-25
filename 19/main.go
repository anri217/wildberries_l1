package main

import "fmt"

func reverseString(s string) (result string) {
	for _, val := range s {
		result = string(val) + result
	}
	return
}

func main() {
	str := "главрыба"
	fmt.Println("Reversed string:", reverseString(str))
}
