package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	square := square{sideLength: 10}
	triangle := triangle{base: 10, height: 15}

	printArea(square)
	printArea(triangle)

}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func printArea(s shape) {
	fmt.Println(s)
	fmt.Println(s.getArea())
}
