package main

import (
	"fmt"
)

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I'm a gator")
}

func main() {
	var g1 gator

	g1.greeting()
}
