package main

import "fmt"

type street struct {
	number int
	road string
	code string
}
type house struct {
	street
	rooms int
	room string
	furniture string
}
func main()  {
nicky := house{
		street:street{
		number: 18,
		road: "Nightingale Gardens",
		code: "CF32 1GB",
	},
	rooms: 8,
	room: "Living Room",
	furniture: "Sofa, TV and sideboard",
}
fmt.Println(nicky)
}