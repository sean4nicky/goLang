package main

import (
	"fmt"
	"encoding/json"
)
/*
Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package
*/
type juser struct {
	First string
	Age   int
}

func main() {
	u1 := juser{
		First: "James",
		Age:   32,
	}

	u2 := juser{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := juser{
		First: "M",
		Age:   54,
	}

	users := []juser{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	us, err := json.Marshal(users)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(string(us))


}

