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

	fmt.Println(t)
	fmt.Println(s)

	fmt.Println(t.color)
	fmt.Println(s.luxury)
}
