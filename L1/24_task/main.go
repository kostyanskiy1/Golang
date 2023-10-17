package main

import (
	"fmt"
	"math"
)

type Point struct { //труктурa Point с инкапсулированными параметрами x,y
	x int
	y int
}

func New(x, y int) *Point { // Конструктор
	return &Point{x, y}
}
func (p1 *Point) Dist(p2 *Point) float64 { //нахождениe расстояния между двумя точками
	dx := float64(p1.x - p2.x)
	dy := float64(p1.y - p2.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := New(1, 2)
	p2 := New(5, 7)

	dist := p1.Dist(p2)
	fmt.Println("Расстояние:", dist)
}
