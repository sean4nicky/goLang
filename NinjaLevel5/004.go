package main
import "fmt"
func main(){
	s := struct {
		house string
		rooms []string
		owned bool
		furniture map[string]string

	}{house: "Semi-detached Town House",
		rooms:     []string{"Kitchen", "Bedroom", "Living Room"},
		owned:     true,
		furniture: map[string]string{
		"Kitchen" : "Fridge, Washing Machine, Oven",
		"Bedroom" : "Bed, Wardrobe, Bed Side Tables",
		"Living Room" : "Sofa, Tv, Cabinet",

		},
	}
	fmt.Println(s.house)
	fmt.Println(s.owned)
	fmt.Println(s.rooms)
	fmt.Println(s.furniture)
}