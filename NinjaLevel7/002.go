package main

import "fmt"

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
