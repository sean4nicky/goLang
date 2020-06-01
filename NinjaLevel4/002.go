package main

import "fmt"

func main()  {
	s := []int{10,20,30,40,50,60,70,80,90,100}
	for i, v := range s {
		fmt.Printf("Index=: %v\t Values=: %d\t Type=: %T\n",i, v, s)
	}
}