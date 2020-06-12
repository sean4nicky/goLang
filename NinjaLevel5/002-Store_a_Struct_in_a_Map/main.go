package main

import (
	"fmt"
)
/*
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
*/

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