package main

import "fmt"

func main() {

	n := 18
	if n >= 60 {
		fmt.Println("old age")
	} else if n >= 18 && n <=60 {
		fmt.Println("major")
	} else {
	fmt.Println("minor")
	}
}