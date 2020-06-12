package main

import "fmt"
/*
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
 */
func main()  {
	favSport := string("Darts")
	switch favSport {
	case "Snooker":
		fmt.Println("my fav sport is: Snooker")
	case "Football":
		fmt.Println("my fav sport is: Football" )
	default:
		fmt.Println("I don't like sport!")

	}
}