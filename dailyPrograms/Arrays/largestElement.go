package main

import "fmt"

func main() {

	array1 := []int{1, 24, 34, 46, 87, 8, 87, 775}
	temp := 0
	for i := 0; i <= len(array1); i++ {
		for j := i + 1; j <= len(array1)-1; j++ {
			if array1[i] > array1[j] {
				temp = array1[i]
				array1[i] = array1[j]
				array1[j] = temp
			}
		}
	}
	fmt.Println("sorted array:",array1)
	element := array1[len(array1)-1]
	fmt.Println("largest element :",element)
	fmt.Println("smallest element in array :",array1[0])
}