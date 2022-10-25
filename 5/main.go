package main

import (
	"fmt"
	"time"
)

func read(ch *chan int) {
	for v := range *ch {
		fmt.Println(v)
	}
}

func send(ch *chan int) {
	i := 0
	for {
		*ch <- i
		i++
	}
}

func main() {
	n := 5
	ch := make(chan int)
	go read(&ch)
	go send(&ch)
	sec_count := n * int(time.Second)
	time.Sleep(time.Duration(sec_count))
	close(ch)
}
