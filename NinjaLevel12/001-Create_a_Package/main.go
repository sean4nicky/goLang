package main

import (
	"fmt"
	"github.com/sean4nicky/goLang/NinjaLevel12/001-Create_a_Package/dog"
)

type bread struct {
	name string
	age int
}
func main() {
	d1 := bread{name: "rover", age: 5}
	fmt.Printf("%v is %v in Human Years and %v in Dog Yers", d1.name, d1.age, dog.Years(d1.age))

}
