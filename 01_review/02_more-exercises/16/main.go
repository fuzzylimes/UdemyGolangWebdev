package main

import (
	"fmt"
)

type gator int

type flamingo bool

func (g gator) greeting() {
	fmt.Println("Hello, I'm a gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

type swampCreature interface {
	greeting()
}

func bayou(s swampCreature) {
	s.greeting()
}

func main() {
	var g1 gator
	var f1 flamingo
	bayou(g1)
	bayou(f1)
}
