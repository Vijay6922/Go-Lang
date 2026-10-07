package main

import "fmt"

// func print[T int | string](num [] T)  -- accepts both int and string 
// func print[T comparable](num [] T) -- can also be used for all types

type stack[T any] struct {
	elements []T
}

func print[T interface{}](num []T) {
	for _, k := range num {
		fmt.Print(k,"\t")
	}
}

func main() {

	nums := []int {1,2,3,4}
	nums1 := [] string{"one","two"}
	print(nums)
	fmt.Println()
	print(nums1)
	fmt.Println()
	mystack := stack[string]{
		elements: []string{"vijay"},
	}
	fmt.Println(mystack)
	
	mystackint := stack[int]{
		elements: []int{1,3,4,4},
	}
	fmt.Println(mystackint)
}