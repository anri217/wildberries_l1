package main

import "fmt"

// parent
type Human struct {
}

// child
type Action struct {
	Human
}

// parent method
func (h *Human) helloWorld() {
	fmt.Println("Hello, world!")
}

// parent method
func (h *Human) SayHello(name *string) {
	fmt.Println("Hello, " + *name + "!")
}

// Action has all Human methods
func main() {
	action := Action{}
	action.helloWorld()
	name := "Max"
	action.SayHello(&name)
}
