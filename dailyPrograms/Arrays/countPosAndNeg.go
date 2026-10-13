package main

import (
	"fmt"
)
func main() {

	array1 :=[]int{1,24,5,-4,3,0,-0}
	posCount := 0
	negCount :=0
	for _,i:= range array1 {
		if i > 0 {
			posCount++
		} else{
			negCount++
		}
	}
	fmt.Println("postive numbers : ",posCount)
	fmt.Println("negative numbers : ",negCount)

}