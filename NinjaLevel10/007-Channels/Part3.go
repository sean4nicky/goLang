package main

import "fmt"

//write a program that
//○ launches 10 goroutines
//■ each goroutine adds 10 numbers to a channel
//○ pull the numbers off the channel and print them
//var wg sync.WaitGroup

/*func main() {
	var gr = 10
	for i := 1; i <= gr; i++{
		fmt.Printf("GoRoutines: %v\n", i)
		for j := 1; j <= 10; j++{
			fmt.Printf("Channel: %v\n", j)
		}
	}

}*/
type result struct {
	rou []int
	cha chan int
}


func rou(){
	for i := 1; i <= 10; i++ {
	}
}
func cha(c chan result){
	for j := 1; j <= 10; j++{
		c <- result(int(j))
	}
}
func main() {
	c:= make(chan result)
	//r:= make()
	go rou()
	go cha(c)
	fmt.Print()
	fmt.Print(<-c)

}