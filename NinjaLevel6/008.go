package main

import "fmt"

//import "fmt"

func main() {
	a := foo()
		fmt.Printf("%T\n", a())
	a() // only puts 64 int memory
	fmt.Println("you called foo", a())
	}

func foo()func()int{
	fmt.Println("before return func")
	return func()int{
		fmt.Println("after return 64")
		return 64
	}
}
