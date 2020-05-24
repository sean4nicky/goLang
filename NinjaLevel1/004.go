package main

import "fmt"

type number int
var x number
func main () {
	fmt.Println("Initial VALUE of x: = ", x)
	fmt.Print("TYPE of x: = ")
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println("This is the new VALUE of x: = ", x)
}