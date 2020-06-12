package main

import "fmt"
/*
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
*/
func main()  {
	dob := 1970
	for dob <= 2020{
		fmt.Println(dob)
		dob++
	}
}
