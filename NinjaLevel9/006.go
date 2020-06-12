package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This is my Architecture", runtime.GOARCH)
	fmt.Println("This is my OS", runtime.GOOS)
}
