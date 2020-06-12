package main

import "fmt"
/*
Build and use an anonymous func
*/
func main(){
	func() {
		a := 10 + 10
		fmt.Println(a)
}()
}
