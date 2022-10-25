package main

import "fmt"

type KeyValue struct {
	Key   string
	Value string
}

// write DATA for map
func writeInChan(ch chan KeyValue) {
	ch <- KeyValue{"1", "2"}
	ch <- KeyValue{"3", "4"}
	ch <- KeyValue{"5", "6"}
	close(ch)
}

func main() {
	ch := make(chan KeyValue)
	mp := make(map[string]string)

	go writeInChan(ch)
	// for each key and value - add this pair to map
	for val := range ch {
		mp[val.Key] = val.Value
	}
	// print map
	fmt.Println(mp)
}
