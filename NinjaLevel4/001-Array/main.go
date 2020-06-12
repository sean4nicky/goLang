package main

import "fmt"
/*
Using a COMPOSITE LITERAL:
create an ARRAY which holds 5 VALUES of TYPE int
assign VALUES to each index position.
Range over the array and print the values out.
Using format printing
print out the TYPE of the array
*/
func main() {
	a := [5]int{10, 20, 30, 40, 50}
	for i, v := range a {
		fmt.Printf("Index=: %v\t Value=: %d\t Type=: %T\n", i, v, a)
	}
	}