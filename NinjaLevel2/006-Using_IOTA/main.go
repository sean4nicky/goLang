package main

import "fmt"
/*
Create a variable of type string using a raw string literal. Print it.
*/
const (
	year1 = iota + 2020
	year2 = iota + 2020
	year3 = iota + 2020
	year4 = iota + 2020

)

func main()  {
fmt.Println(year1)
fmt.Println(year2)
fmt.Println(year3)
fmt.Println(year4)
}
