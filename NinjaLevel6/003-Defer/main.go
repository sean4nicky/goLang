package main

import "fmt"
/*
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/
func main() {
	ii := []int{1,2,3,4,5,6,7,8,9}
	xi := foo(ii...)
	defer fmt.Println(xi)

	ii2 := []int{11,12,13,14,15,16,17,18,19}
	xi2 := bar(ii2)
	fmt.Println(xi2)
}
func foo(i ...int) int{
	total := 0
for _, v := range i {
	total += v
}
return total
}
func bar(i []int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}