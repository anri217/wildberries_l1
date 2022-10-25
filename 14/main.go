package main

import (
	"fmt"
	"reflect"
)

func main() {
	var val interface{}

	val = 1.1

	reflect_val := reflect.ValueOf(val)

	switch reflect_val.Kind() {
	case reflect.String:
		fmt.Println("type: String")
	case reflect.Int:
		fmt.Println("type: Int")
	case reflect.Bool:
		fmt.Println("type: Bool")
	case reflect.Float64:
		fmt.Println("type: Float64")
	case reflect.Chan:
		fmt.Println("type: Chan")
	}
}
