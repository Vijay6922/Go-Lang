package main

import "fmt"

func main () {

	num := 5 
	cube :=1
	for i:=0; i<3;i++ {
		cube *= num
	}
	fmt.Println("cube of number is :",cube)
}