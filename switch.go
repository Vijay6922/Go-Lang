package main

import (
	"fmt"
	"time"
)

func main() {

	switch time.Now().Weekday(){
		case time.Sunday, time.Saturday : {
			fmt.Println("weekend")
		}
		default :{
			fmt.Println("weekday")
		}
	}

	i := 2012

	switch {
	case i%400 ==0 : {
		fmt.Println("leap year")
	}
	case i%100==0 : {
		fmt.Println("not leap year")
	}
	case i%4==0 : {
		fmt.Println("leap year")
	}
	default :{
		fmt.Println("not a leap year")
	}
	}


}
