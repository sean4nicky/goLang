package main

import "fmt"
/*
Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop
*/
func main() {
	m := map[string]string{
		"James Bond":      "Shaken, not stirred, Martinis, Women",
		"Moneypenny Miss": "James Bond, Literature, Computer Science",
		"No Dr":           "Being evil, Ice cream, Sunsets",
	}
	fmt.Println(m)

	m["Q"] = "Computers, Guns, Secrets"

	for i, v := range m {
		fmt.Printf("%v\t\t\t%v\n", i, v)
	}
}
