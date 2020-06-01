package main

import (
	"fmt"
)


type person struct {
	first_name string
	last_name string
	favorite_ice_cream_flavors []string
}
func main() {

	p1 := person{
		first_name:                 "Nicky ",
		last_name:                  "Murphy",
		favorite_ice_cream_flavors: []string{"toffe", "bannana", "coffee"},
	}
	p2 := person{
		first_name:                 "Ryan ",
		last_name:                  "seanMurphy",
		favorite_ice_cream_flavors: []string{"chocolate", "strawbery", "vanilla"},
	}
	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}
	for _, v := range m {
		fmt.Println(v.first_name)
		for i, val := range v.favorite_ice_cream_flavors {
			fmt.Println(i, val)

		}
	}
}