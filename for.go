package main

import "fmt"

func main() {

	//while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i+1
	}

	//infite loop 
	// for {
	// 	fmt.Println("hi")
	// }

	//classic for loop
	for j :=10; j>=1; j-- {
		fmt.Println(j)
	}

	//range 
	for k := range 30 {
		fmt.Print(k)
	}
}