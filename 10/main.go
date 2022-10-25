package main

import "fmt"

func main() {
	// temps
	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// map store [] floats
	group_temps := make(map[int][]float64)

	for _, value := range array {
		// round calue by 10
		i := int(value) - (int(value) % 10)
		group_temps[i] = append(group_temps[i], value)
	}

	fmt.Println(group_temps)
}
