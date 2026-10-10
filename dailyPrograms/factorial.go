package main

import "fmt"

func main () {

	num := 5
	fact :=1
	for i:=num;i>=1;i--{
		fact *= i
	}
	fmt.Println("Factorial of the number is :",fact)
}