package main

import (
	"fmt"
	"slices"
)

func uniqueElements(a []int)[]int {

	var result []int
	for i:=0;i<len(a);i++ {
		if !contains(result,a[i]) {
			result=append(result,a[i])
		}
	}
	return result
}

func contains(a []int,value int)bool{

	for i:=0;i<len(a);i++ {
		if a[i] ==value{
			return true
		}
	}
	return false
}


func main () {
	count:=0
	array1:=[]int{1,2,5,65,6,5,7,5}
	array2:=[]int{}
	for i:=0;i<=len(array1);i++ {
		for j:=i+1;j<=len(array1)-1;j++ {
		if array1[i]==array1[j] {
			count=count+1
			array2=append(array2, array1[i])
		}
	}
	}
	fmt.Println("duplicate element",array2)
	fmt.Println("repeated for :",count)

	//need to sort fisrt to remove the duplicate elements 
	slices.Sort(array1)
	result := slices.Compact(array1)
	fmt.Println("unique array : ",result)

	array3:=[]int{11,188,8,186,86,6,8,6,68,68,11}
	uniq := uniqueElements(array3)
	fmt.Println("removed duplicate elements using loops ",uniq)

}