package main

import (
	"fmt"
	"math/rand"
)

// we use const string because in golang strings can be indexed which indexes its bytes
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	fmt.Println(RandStringBytes(1 << 10)[:100])
}
