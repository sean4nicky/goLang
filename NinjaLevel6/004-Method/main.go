package main

import "fmt"
/*
Create a user defined struct with
the identifier “person”
the fields: first, last, age,
attach a method to type person with
the identifier “speak”
the method should have the person say their name and age
create a value of type person
call the method from the value of type person
*/
type person struct {
	first string
	last string
	age int
}
type robot struct {
	person
	OS string
	AI bool
}

func (p person) speak()  {
	fmt.Println("Hello my name is", p.first, p.last, "and I am ", p.age, "Years old"  )
}
type human interface {
	speak()
}

func main() {
	per1 := person {
		first: "sean",
		last: "Murphy",
		age: 49,
	}
	fmt.Println(per1.first, per1.last, per1.age)
	Per2 := robot {
		person : person{
			first: "Brian",
			last: "Board",
			age: 16,
		},
		OS: "Android",
		AI: true,
	}
fmt.Println(Per2)


per1.speak()
Per2.speak()
}
