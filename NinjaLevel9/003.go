package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int
var wait sync.WaitGroup

func main() {
	fmt.Println("CPU count Start: \t", runtime.NumCPU())
	fmt.Println("GoRoutines count Start: \t",runtime.NumGoroutine())

	var gr = 30
	wait.Add(gr)
	for i := 0; i < gr; i++ {
		go foo()
		wait.Done()
	}
	wait.Wait()
	fmt.Println("GoRoutines count END: \t",runtime.NumGoroutine())

}



func foo()  {
		val := count
		runtime.Gosched()
		val++
		count = val
		fmt.Println("Value of Count: \t",count)
	}
