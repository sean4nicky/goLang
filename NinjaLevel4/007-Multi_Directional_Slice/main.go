package main

import "fmt"
/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."
Range over the records, then range over the data in each record.
*/
func main() {
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	s := [][]string{a, b}
	for i, v := range s {
		fmt.Println(i, v)
	}  /*for i, v := range s {
	fmt.Println("record: ", i)
	for j, val := range s {
		fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
	}
}/*
	for i, v := range a {
		fmt.Print(i, " : ", v, "   ")
	}
	fmt.Println(" ")
	for i, v := range b {
		fmt.Print(i, " : ", v, "   ")
	}*/
}