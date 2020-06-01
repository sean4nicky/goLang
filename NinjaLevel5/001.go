package main

import (
	"fmt"

)

type person struct {
	first_name string
	last_name string
	favorite_ice_cream_flavors []string
}
func main()  {

	p1 := person{
		first_name: "Nicky ",
		last_name: "Murphy",
		favorite_ice_cream_flavors : []string{"toffe", "bannana", "coffee"},
		}
	p2 := person{
		first_name: "Ryan ",
		last_name: "Murphy",
		favorite_ice_cream_flavors: []string{"chocolate", "strawbery", "vanilla"},
	}
	fmt.Print(p1.first_name)
	fmt.Println(p1.last_name)
	for i, v := range p1.favorite_ice_cream_flavors{
		fmt.Println(i, v)
	}
	fmt.Print(p2.first_name)
	fmt.Println(p2.last_name)
	for i, v := range p2.favorite_ice_cream_flavors{
		fmt.Println(i, v)
	}
}