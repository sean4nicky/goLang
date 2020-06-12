package main

import "fmt"
/*
Create a variable of type string using a raw string literal. Print it.
*/
func main() {
	fmt.Println("this is a string")
	foo := `this is a example of a
raw "string()" literal
in can be typed over 
many lines and include
^$!\@ in your typeface`
	fmt.Println(foo)
}