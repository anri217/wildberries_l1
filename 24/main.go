package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// find distance by pifagor phormula
func Distance(point1, point2 *Point) float64 {
	var dx float64
	var dy float64
	if point2.x > point1.x {
		dx = math.Pow(point2.x-point1.x, 2)
	} else {
		dx = math.Pow(point1.x-point2.x, 2)
	}
	if point2.y > point1.y {
		dy = math.Pow(point2.y-point1.y, 2)
	} else {
		dy = math.Pow(point1.y-point2.y, 2)
	}
	return math.Sqrt(dx + dy)
}

func main() {
	point1 := NewPoint(2, 2)
	point2 := NewPoint(1, 1)
	fmt.Println("Distance between X and Y: ", Distance(point1, point2))
}
