package main

import "fmt"

func main () {

	num := 12345

	count := 0

	for {
		if num !=0 {
			count = count+1
			num = num/10
		} else {
			break
		}
	}
	fmt.Println(count)
}
