package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var coun int
var wai sync.WaitGroup
var mu sync.Mutex
var at atomic.Value

func main() {
	fmt.Println("CPU count Start: \t", runtime.NumCPU())
	fmt.Println("GoRoutines count Start: \t",runtime.NumGoroutine())

	var gr = 30
	wai.Add(gr)
	at.Store(foo3)
	for i := 0; i < gr; i++ {
		go foo3()
		wai.Done()
	}
	wai.Wait()

	fmt.Println("GoRoutines count END: \t",runtime.NumGoroutine())

}



func foo3()  {
		mu.Lock()
		val := coun
		val++
		coun = val
		fmt.Println("Value of Count: \t",coun)
		at.Load()
		mu.Unlock()
	}
