package main

import "fmt"

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
