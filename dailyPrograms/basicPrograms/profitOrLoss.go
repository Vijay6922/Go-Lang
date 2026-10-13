package main

import "fmt"

func main() {

	buyingCost := 30
	sellingCost := 50
	if buyingCost > sellingCost {
		fmt.Println("loss")
	} else {
		fmt.Println("profit")
	}
}