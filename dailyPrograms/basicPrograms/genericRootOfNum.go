package main

import "fmt"
//add the digits of the numbers until the sum is less than 10
func main () {

	num:= 2345
	sum := 0

	for {

		r:= num % 10
		sum = sum+r
		num = num / 10

		if sum > 10 {
			num = sum
			sum = 0
		} else {
			break
		}
	}
	fmt.Println("Generic root of the number is :",sum)
}