package main

import (
	"fmt"
	"time"
)

func sleep(secs *int) {
	<-time.After(time.Second * time.Duration(*secs))
}

func main() {
	secs := 10
	fmt.Println(time.Now().Unix())
	sleep(&secs)
	fmt.Println(time.Now().Unix())
}
