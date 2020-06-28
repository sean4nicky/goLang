//Create a custom error message using “fmt.Errorf”.
package main
import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
		if err != nil {
			fmt.Errorf("toJason caused an error: %v", err)
	}
	fmt.Println(string(bs))
	return
//return stops below from being printed.
fmt.Println(bs)
fmt.Println(&err)
fmt.Printf("%#v", bs)
fmt.Printf("%#v",err)
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		fmt.Errorf("Marshal caused an error: %v", err)
	}
	return bs, err
}