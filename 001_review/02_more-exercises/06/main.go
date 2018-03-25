package main

import (
	"fmt"
)

type person struct {
	fName   string
	lName   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintf("%v is walking.", p.fName)
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

	s := p1.walk()
	fmt.Println(s)
}
