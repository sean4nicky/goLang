package main

import "fmt"
/*
create a func with the identifier foo that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in
create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in
*/
func main() {
	ii := []int{1,2,3,4,5,6,7,8,9}
	xi := foo(ii...)
fmt.Println(xi)

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