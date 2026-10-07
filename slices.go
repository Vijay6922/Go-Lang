package main

import "fmt"

func main() {

	// var nums = [] int{12,34,435,54,23}

	// fmt.Println(nums)

	nums1 := [] int {1,2,3,34,4,0,2,1,32,3}
	// fmt.Println(nums1)
	// fmt.Println(len(nums1))
	temp:=0;


	for i := 0; i< len(nums1); i++{
		for j:=0;j<len(nums1)-i-1;j++{
	if nums1[j]>nums1[j+1] {
		temp =nums1[j]
		nums1[j]=nums1[j+1]
		nums1[j+1] = temp
	
	}
}
}
fmt.Println(nums1)
nums3 := len(nums1)-2
fmt.Println(nums1[nums3])

nums12 := [] int {32,3}
temp1 := 0;
for i :=0;i<=len(nums12)-1;i++ {
	temp1 += nums12[i]
}
fmt.Println(temp1)
}

