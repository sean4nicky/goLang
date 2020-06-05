package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width float64
}
type circle struct {
	radius float64
}

func (C circle) area() float64 {
	fmt.Print("Area of Circle with a RADIUS of ", C.radius, "  =  :  ")
	return math.Pi * C.radius * C.radius
}
func (S square) area() float64 {
	fmt.Print("Area of Square ", S.length, " x ", S.width, "  =  :  ")
return S.length * S.width
}
func info (D shape) {
	      fmt.Println("The area is ", D.area() )
}

type shape interface {
	area()float64
}


func main() {
	sq := square{
		length: 10,
		width: 10,
	}
	cr := circle{
		radius: 3,
	}
	fmt.Println(sq.area())
	fmt.Println(cr.area())
	info(sq)
	info(cr)

}
