package main

import "fmt"

func main()  {
	favsport := string("Darts")
	switch favsport {
	case "Snooker":
		fmt.Println("my fav sport is: Snooker")
	case "Football":
		fmt.Println("my fav sport is: Football" )
	default:
		fmt.Println("I don't like sport!")

	}
}