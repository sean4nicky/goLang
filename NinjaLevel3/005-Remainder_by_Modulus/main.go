package main

import "fmt"
/*
Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
 */
func main() {
	for i := 10; i <= 100; i++{
		times := i/4
		fmt.Printf("%v /4 = %v Remainder %v\n",i, times, i%4)
		}

	}
