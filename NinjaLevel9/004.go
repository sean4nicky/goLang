package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counts int
var waits sync.WaitGroup
var mut sync.Mutex

func main() {
	fmt.Println("CPU count Start: \t", runtime.NumCPU())
	fmt.Println("GoRoutines count Start: \t", runtime.NumGoroutine())

	var gr = 30
	waits.Add(gr)
	for i := 0; i < gr; i++ {
		go foo2()
		waits.Done()
	}
	waits.Wait()

	fmt.Println("GoRoutines count END: \t", runtime.NumGoroutine())

}

func foo2() {
	mut.Lock()
	val := counts
	val++
	counts = val
	fmt.Println("Value of Count: \t", counts)
	mut.Unlock()
}
