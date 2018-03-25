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

func (t truck) transportationDevice() string {
	return "I'm a truck. I haul crap."
}

func (s sedan) transportationDevice() string {
	return "I'm a car. I haul human garbage."
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

	fmt.Println(t.transportationDevice())
	fmt.Println(s.transportationDevice())
}
