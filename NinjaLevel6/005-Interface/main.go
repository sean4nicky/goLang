package main

import (
	"fmt"
	"math"
)
/*
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/
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
