package main

import "fmt"

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