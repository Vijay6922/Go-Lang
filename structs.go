package main

import (
	"fmt"
	"time"
)

//define struct
type order struct {
	id int
	name string 
	amount float32
	createdAt time.Time
	Car
}

type Car struct {
	name string
	origin string
}

func (o *order) change(name string) string {
	o.name = name
	return name
}

func (o *order) interestPerAnnum(amount float32, rate float32) float32 {
	simpleInterest := (amount*rate*1)/100
	return simpleInterest
}

func main () {

	myOrder := order{
		id : 1,
		name : "vijay",
		amount:12324.00,
		createdAt: time.Now(),
	}
	myOrder.createdAt= time.Time{}

	mycars := Car {
		name: "bmw",
		origin: "germany",
	}

	myOrder2 := order{
		id : 1,
		name : "kumar",
		amount:123.00,
		createdAt: time.Now(),
		Car : mycars,      
	}
	myOrder2.Car.name="ferrari"
	fmt.Println(myOrder)
	fmt.Println(myOrder2)
	myOrder2.change("chintu")
	fmt.Println(myOrder2.name)
	fmt.Println(myOrder2.interestPerAnnum(350000,9))

	//dynamic struct 
	language := struct {
		name string
		isGood bool
	} {"telugu",true}
	fmt.Println(language)

	

}
