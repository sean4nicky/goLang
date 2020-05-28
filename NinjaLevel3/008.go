package main

import (
	"fmt"
	"time"
)

func main()  {
	timer1 := time.NewTimer(time.Second * 5)
	timer2 := time.NewTimer(time.Second * 10)
	timer3 := time.NewTimer(time.Second* 15)
	for i:=0;i<100;i++{
		switch  {
		case i == 10:
			fmt.Println("you have reached 10")
		case i == 50:
			<-timer1.C
			fmt.Println("you have reached 50")
		case i == 90:
			<-timer2.C
			fmt.Println("you are almost at 100")
		//default:
		<-timer3.C
			fmt.Println("you have made it to 100")
		}

	}
}