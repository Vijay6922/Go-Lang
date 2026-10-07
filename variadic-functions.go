package main

import "fmt"

func sum(add ...int) int {
	total := 0
	for i,nums := range add {
		total = total + nums
		fmt.Println(i,nums)
	}
	return total
}

func main() {

	nums := []int{1, 3,9,9}

	result := sum(nums...)
	fmt.Println(result)

}