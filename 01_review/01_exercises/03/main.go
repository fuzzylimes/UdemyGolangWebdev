package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
}

func (p person) pSpeak() string {
	return "Hello, friend."
}

func (sa secretAgent) saSpeak() string {
	return "That's my bottom."
}

func main() {
	p := person{"Bob", "Joe"}
	sa := secretAgent{
		person{
			"Jill",
			"Jane",
		},
	}

	fmt.Println(p.lname)
	fmt.Println(p.pSpeak())

	fmt.Println(sa.fname)
	fmt.Println(sa.saSpeak())
	fmt.Println(sa.person.pSpeak())
}
