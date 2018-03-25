package main

import (
	"fmt"
)

type gator int

func main() {
	var g1 gator
	var x int
	g1 = 7
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = int(g1)
	fmt.Println(x)
}
