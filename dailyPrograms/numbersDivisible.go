package main 

import "fmt"

func main() {

	limit := 200

	nums := []int{2,3,4,5,6,7,8,9}

	for i:=0;i<len(nums);i++ {
		fmt.Printf("Numbers divisible by %d are:\n", nums[i])
		for j:=1;j<=limit;j++ {
			if j%nums[i]==0 { //here nums[i] is the divisor and j is the number being checked for divisibility
				fmt.Print(j, " ")
			}
		}
		fmt.Println()
	}
}