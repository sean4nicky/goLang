package main

import (
	"fmt"
	"runtime"
	"sync"
)
/*
Using goroutines, create an incrementer program
have a variable to hold the incrementer value
launch a bunch of goroutines
each goroutine should
read the incrementer value
store it in a new variable
yield the processor with runtime.Gosched()
increment the new variable
write the value in the new variable back to the incrementer variable
use waitgroups to wait for all of your goroutines to finish
the above will create a race condition.
Prove that it is a race condition by using the -race flag
if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
*/
var count int
var wait sync.WaitGroup

func main() {
	fmt.Println("CPU count Start: \t", runtime.NumCPU())
	fmt.Println("GoRoutines count Start: \t", runtime.NumGoroutine())

	var gr = 30
	wait.Add(gr)
	for i := 0; i < gr; i++ {
		go foo()
		wait.Done()
	}
	wait.Wait()
	fmt.Println("GoRoutines count END: \t", runtime.NumGoroutine())

}

func foo() {
	val := count
	runtime.Gosched()
	val++
	count = val
	fmt.Println("Value of Count: \t", count)
}
