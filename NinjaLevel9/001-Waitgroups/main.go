package main

import (
	"fmt"
	"runtime"
	"sync"
)
/*
in addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
*/
var counter int
var wg sync.WaitGroup

func main() {
	fmt.Println("CPU count Start: \t", runtime.NumCPU())
	fmt.Println("GoRoutines count Start: \t", runtime.NumGoroutine())

	wg.Add(2)

	go boo()
	go bar()

	wg.Wait()

	fmt.Println("GoRoutines count End: \t", runtime.NumGoroutine())
	fmt.Println("the end : :)")

}

func boo() {
		fmt.Println("count of boo: \t", counter)
		wg.Done()
	}


func bar() {
	fmt.Println("count of bar: \t", counter)
	wg.Done()
	}
