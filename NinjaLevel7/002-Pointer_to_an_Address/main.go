package main

import "fmt"
/*
create a person struct
create a func called “changeMe” with a *person as a parameter
change the value stored at the *person address
important
to dereference a struct, use (*value).field
p1.first
(*p1).first
“As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
https://golang.org/ref/spec#Selectors
create a value of type person
print out the value
in func main
call “changeMe”
in func main
print out the value
*/
type person struct {
	first string
	last string
	age int
}

func changeMe(p *person)  {
		(*p).first = "Sean"
		(*p).last = "Murphy"
		(*p).age = 49
}

func main() {

	p1 := person {
		first: "Howard",
		last: "Jones",
		age: 29,
	}
	fmt.Println(p1)

	changeMe(&p1)

	fmt.Printf("%T\n", p1)
	fmt.Println(p1)
}
