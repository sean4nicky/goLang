package main

import "fmt"

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
	}
}