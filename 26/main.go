package main

import "fmt"

func isUniqueLetters(str string) bool {
	// for each i elem we check all other elems in arr
	for i := 0; i < len(str)-1; i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(isUniqueLetters("abcd"))
	fmt.Println(isUniqueLetters("abCdefAaf"))
	fmt.Println(isUniqueLetters("aabcd"))
}
