package main
import "fmt"
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