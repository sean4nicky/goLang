package main

import "fmt"

func main() {
	m := map[string]string{
		"James Bond":      "Shaken, not stirred, Martinis, Women",
		"Moneypenny Miss": "James Bond, Literature, Computer Science",
		"No Dr":           "Being evil, Ice cream, Sunsets",
	}
	fmt.Println(m)

	m["Q"] = "Computers, Guns, Secrets"

	fmt.Println(m)

	for i, v := range m {
		fmt.Printf("%v\t\t%v\n", i, v)
	}
	delete(m, "Moneypenny Miss")

	fmt.Println(m)

	for i, v := range m {
		fmt.Printf("%v\t\t%v\n", i, v)
	}
}