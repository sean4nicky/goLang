package main

import (
	"fmt"
	"encoding/json"
)

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

