package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type transportation interface {
	transportationDevice() string
}

func (t truck) transportationDevice() string {
	return "I'm a truck. I haul crap."
}

func (s sedan) transportationDevice() string {
	return "I'm a car. I haul human garbage."
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func main() {
	t := truck{
		vehicle{
			4,
			"black",
		},
		true,
	}

	s := sedan{
		vehicle{
			4,
			"silver",
		},
		false,
	}

	report(t)
	report(s)
}
