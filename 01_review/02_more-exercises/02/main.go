package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Bill": 42,
		"John": 45,
	}

	fmt.Println(m)

	for i := range m {
		fmt.Println(i)
	}

	for i, j := range m {
		fmt.Println(i, j)
	}
}
