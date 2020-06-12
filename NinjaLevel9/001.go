package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mutex = sync.Mutex{}

//var mutex2 = sync.Mutex{}

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

func boo() int {
	i := 0
	for i = 0; i < 10; i++ {
		mutex.Lock()
		counter++
		fmt.Println("count of boo: \t", counter)
		mutex.Unlock()
	}
	wg.Done()
	return counter
}

func bar() int {
	//time.Sleep(time.Nanosecond *1)
	i := 0
	for i = 0; i < 10; i++ {
		mutex.Lock()
		counter += 2
		fmt.Println("count of bar: \t", counter)
		mutex.Unlock()
	}
	wg.Done()
	return counter
}
