package main

import "fmt"

func main() {
	a := [5]int{10, 20, 30, 40, 50}
	for i, v := range a {
		fmt.Printf("Index=: %v\t Value=: %d\t Type=: %T\n", i, v, a)
	}
	}