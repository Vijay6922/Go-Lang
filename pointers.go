package main

import "fmt"

//pointers fucntion
func changeNumber(num *int) int{ 
	*num = *num * *num
	fmt.Println(*num)
	return *num
}

func main () {

	number := 5
	fmt.Println("before changing",number)
	(changeNumber(&number))
	fmt.Println("after changing",number)
}
