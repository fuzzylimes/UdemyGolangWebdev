package main

import (
	"fmt"
	"math"
)

type square struct {
	l float64
	w float64
}

type circle struct {
	r float64
}

func (s square) area() float64 {
	return s.l * s.w
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{10.1, 10.1}
	c := circle{2}
	info(s)
	info(c)
}
