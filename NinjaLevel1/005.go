package main

import "fmt"

type number int
var x number
var y int
func main () {
	fmt.Println("Initial VALUE of x: = ", x)
	fmt.Print("TYPE of x: = ")
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println("This is the new VALUE of x: = ", x)

	y = int(x)
	fmt.Println("The assigned VALUE of y from x: = ", y)
	fmt.Print("TYPE of Y: = ", y)
	fmt.Printf("%T", y)

}