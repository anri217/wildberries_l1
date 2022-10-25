package main

import (
	"fmt"
	"time"
)

func sleep(secs *int) {
	start_time := time.Now().Unix()
	for {
		secs_duration := time.Now().Unix() - start_time
		if secs_duration >= int64(*secs) {
			return
		}
	}
}

func main() {
	secs := 10
	fmt.Println(time.Now().Unix())
	sleep(&secs)
	fmt.Println(time.Now().Unix())
}
