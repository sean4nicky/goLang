package main

import "fmt"
/*
Assign a func to a variable, then call that func
*/

func main() {
	a := func(){
		fmt.Println("you called a")
	}

	a()
}
