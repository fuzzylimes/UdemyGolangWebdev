package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a)

	for i := range a {
		println(i)
	}

	for i, j := range a {
		println(i, j)
	}

}
