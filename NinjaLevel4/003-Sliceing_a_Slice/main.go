package main

import "fmt"
/*
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
*/
func main()  {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a := x[:5]     //Correct solution fmt.Println(x[:5])
	fmt.Println(a)
	b := x[5:]     //Correct solution fmt.Println(x[5:])
	fmt.Println(b)
	c := x[2:][:5] //Correct solution fmt.Println(x[2:7])
	fmt.Println(c)
	d := x[1:][:5] //Correct solution fmt.Println(x[1:6])
	fmt.Println(d)
}