package main
import "fmt"
/*
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results
*/
func main() {
	foo(129)
	bar(17, "Year's old")
}
func foo(n int)int  {
	fmt.Println(n)
	return n
}
func bar(i int, s string)(int,(string)) {
	fmt.Println(i, s)
	return i, s
}