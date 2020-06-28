package main

import "fmt"

//write a program that
//○ puts 100 numbers to a channel
//○ pull the numbers off the channel and print them
func main() {
c100:= gen()
rec(c100)
}

func gen()chan int {
	c100 := make(chan int)
	go func() {
		for i := 0; i < 100; i++{
			c100 <- i+1
		}
		close(c100)
	}()
	return c100
}

func rec(c chan int) {
		for v := range c{
			fmt.Println(v)
		}
}