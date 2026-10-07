package main

import "fmt"

//returns a function
func demo () func() int {
	counter :=0
	return func() int {
		counter=counter+1
		return counter
	}
}
func main () {

	clos := demo()
	fmt.Println(clos())
	fmt.Println(clos())
}

