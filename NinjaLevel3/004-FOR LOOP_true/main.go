package main

import "fmt"
/*
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.
*/
func main()  {
	dob := 1970
	for {
		if dob > 2020 {
			break
		}
		fmt.Println(dob)
		dob++
		}

	}
