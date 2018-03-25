package main

import (
	"fmt"
)

type person struct {
	fName   string
	lName   string
	favFood []string
}

func main() {
	p1 := person{
		fName: "Bob",
		lName: "Joe",
		favFood: []string{
			"pizza",
			"tacos",
			"doritos",
		},
	}
	fmt.Println(p1.favFood)
	for _, i := range p1.favFood {
		fmt.Println(i)
	}
}
