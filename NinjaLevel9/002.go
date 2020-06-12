package main

import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		first: "John",
		last:  "Hammond",
	}

	fmt.Println(p1)
	//saySomething(p1.first, p1.last)
	p1.speak()
	saySomething(&p1)
}
