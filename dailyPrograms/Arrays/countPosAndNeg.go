package main

import (
	"fmt"
)
func main() {

	array1 :=[]int{1,24,5,-4,3,0,-0}
	posCount := 0
	negCount :=0
	pos:=[]int{}
	neg:=[]int{}
	for _,i:= range array1 {
		if i > 0 {
			posCount++
			pos=append(pos,i)
		}  else if i < 0 {
			negCount++
			neg=append(neg, i)
		}
	}
	fmt.Println("postive numbers : ",posCount)
	fmt.Println("positive numbers :",pos)

	fmt.Println("negative numbers : ",negCount)
	fmt.Println("negative numbers :",neg)

}