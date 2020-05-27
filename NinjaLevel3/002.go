package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v", i)
		fmt.Printf("\n\t%#U\n\t%#U\n\t%#U\n", i, i, i)

	}
}
