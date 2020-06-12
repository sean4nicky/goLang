package main

import (
	"fmt"
)
/*
Create TYPED and UNTYPED constants. Print the values of the constants.
*/
const (
	foo = true
	bar bool = false
)
func main() {
		fmt.Println("bar is always FALSE as a set constant can not be changed.")

		fmt.Println("foo is Always TRUE as it is a constant and can't be changed.")
}
