package main

import "fmt"
/*
Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable:
*/
func main() {
	f :=sum(1)
fmt.Println(f)
g := sum(5)
fmt.Println(g)

}
func sum(s int)int {
	a := s
	for a < 10 {
		fmt.Println(a)
		a++
	}
	return a
}