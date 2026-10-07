package main

import "fmt"

type status int 

const (
	Recieved status = iota 
	confimed 
	prepared 
	deliverd
)

func change(status status) {
	fmt.Println("changed",status)
}

type orderstatus string 

const (
	Booked orderstatus = "Booked"
	done = "done"
)

func statuses(status orderstatus) {
	fmt.Println("changed",status)
}
func main() {
	change(Recieved)
	change(confimed)
	change(prepared)
	change(deliverd)

	statuses(Booked)
}
