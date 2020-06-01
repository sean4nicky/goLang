package main

import "fmt"

type vehicle struct {
	doors int
	colour string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}
func main()  {
	mac := truck{
		vehicle : vehicle {
			doors: 2,
			colour: "red",
		},
		fourWheel: true,
	}
	sed := sedan{
		vehicle : vehicle {
			doors:  4,
			colour: "Blue",
		},
		luxury: true,
	}
	fmt.Println(mac)
	fmt.Println(mac.doors)
	fmt.Println(sed)
	fmt.Println(sed.colour)
}