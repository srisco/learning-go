package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{
		sideLength: 2.0,
	}

	printArea(s)

	t := triangle{
		height: 3.0,
		base:   1.5,
	}
	printArea(t)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}
