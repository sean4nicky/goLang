package main

import "fmt"

func main()  {
	a := 100
	if a == 20 {
		fmt.Println("a is 20")
	} else if a == 10 {
		fmt.Println("a is 10")
	} else {
		println("Count aborted")
	}
}