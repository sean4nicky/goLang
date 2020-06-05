package main

import "fmt"


func main()  {
	one(90)
	two(one, 301)
sum := two(one, 210)
fmt.Println(sum)
}

func two(t func(x int)int, v int)int{
	for i := v; i <= 500; i++ {
		if i == 100{
			fmt.Println("You are in two")
		} else  if i == 200{
			fmt.Println("you are well in two")
		} else if  i == 300 {
			fmt.Println(" you have reached your goal")
		} else {
			fmt.Println("something went wrong")
		}
		return i
	}
	return t(v)
}

func one(x int)int {
	for x < 100 {
		switch {
		case x > 100:
			break
		case x < 100:
			fmt.Println(x)
			x++
		}
	}
	return x
}
