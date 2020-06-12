package main

import "fmt"

/*
Write down what these print:
fmt.Println(true && true)
fmt.Println(true && false)
fmt.Println(true || true)
fmt.Println(true || false)
fmt.Println(!true)
*/
func main() {
	fmt.Println(true && true) // TRUE
	fmt.Println(true && false)// FALSE
	fmt.Println(true || true) // TRUE
	fmt.Println(true || false)// TRUE
	fmt.Println(!true)        // FALSE
}
