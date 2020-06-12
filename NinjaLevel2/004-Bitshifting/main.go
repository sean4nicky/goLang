package main

import "fmt"
/*
Write a program that
assigns an int to a variable
prints that int in decimal, binary, and hex
shifts the bits of that int over 1 position to the left, and assigns that to a variable
prints that variable in decimal, binary, and hex
*/
var x int = 365
var y = x << 1
func main() {
	fmt.Println("X = :", x)
	fmt.Print("The Decimal value of X = : ")
	fmt.Printf("%d\n", x)
	fmt.Print("The Binary value of X = : ")
	fmt.Printf("%b\n", x)
	fmt.Print("The Hex value of X = : ")
	fmt.Printf("%x", x)

	fmt.Println("X = :", y)
	fmt.Print("The Decimal value of y shifted by 1 = : ")
	fmt.Printf("%d\n", y)
	fmt.Print("The Binary value of y shifted by 1 = : ")
	fmt.Printf("%b\n", y)
	fmt.Print("The Hex value of y shifted by 1 = : ")
	fmt.Printf("%x", y)
}

