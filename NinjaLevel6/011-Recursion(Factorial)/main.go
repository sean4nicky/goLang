package main

import "fmt"
// recursion example of factorial
func main() {
	j := fac(5)
	fmt.Print(j)
}

func fac(n int)int  {
	total := 1
	for; n > 0; n-- {
		total *= n
	}
	return total
}