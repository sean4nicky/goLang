package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++{
		times := i/4
		fmt.Printf("%v /4 = %v Remainder %v\n",i, times, i%4)
		}

	}
