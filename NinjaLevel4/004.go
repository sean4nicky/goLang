package main

import "fmt"

func main()  {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a := append(x, 52)  //x = append(x, 52)
	fmt.Println(a)  // fmt.Println(x)
	b:= append(a, 53, 54,55) //x = append(x, 53, 54, 55)
	fmt.Println(b)  // fmt.Println(x)
	y := []int{56,57,58,59,60}
	c := append(b, y...)  // x = append(x, y...)
	fmt.Println(c)  // fmt.Println(x)
}