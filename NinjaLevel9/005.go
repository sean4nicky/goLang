package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var coun int64
var wai sync.WaitGroup

func main() {
	fmt.Println("CPU count Start: \t", runtime.NumCPU())
	fmt.Println("GoRoutines count Start: \t", runtime.NumGoroutine())

	var gr = 30
	wai.Add(gr)

	for i := 0; i < gr; i++ {
		go foo3()
		wai.Done()
	}
	wai.Wait()

	fmt.Println("GoRoutines count END: \t", runtime.NumGoroutine())

}

func foo3() {
	atomic.AddInt64(&coun, 1)
	fmt.Println("Value of Count: \t", atomic.LoadInt64(&coun))
}
