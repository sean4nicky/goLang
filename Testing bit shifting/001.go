package main

import (
	"fmt"
)

func main() {
	x := 6
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	z := x >> 2
	fmt.Printf("%d\t\t%b\n", z, z)

	a := x << 2
	fmt.Printf("%d\t\t%b", a, a)
}
