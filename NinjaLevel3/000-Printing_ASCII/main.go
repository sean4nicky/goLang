package main
import "fmt"
/*
Print out the ascii value from 33 to 122 so that it shows the string value in accordance to the grid
 */
func main()  {
	for num := 33; num <= 122; num++{
		fmt.Printf("%v\t\t%c\n",num, num)
	}
}
