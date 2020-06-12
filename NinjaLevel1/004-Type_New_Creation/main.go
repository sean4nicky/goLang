package main

import "fmt"

/*For this exercise
Create your own type. Have the underlying type be an int.
create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
in func main
print out the value of the variable “x”
print out the type of the variable “x”
assign 42 to the VARIABLE “x” using the “=” OPERATOR
print out the value of the variable “x”*/

type number int
var x number

func main () {
	fmt.Println("Initial VALUE of x: = ", x)
	fmt.Print("TYPE of x: = ")
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println("This is the new VALUE of x: = ", x)
}