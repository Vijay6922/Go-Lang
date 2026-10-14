package main

import (
	"fmt"
	"slices"
)

func revString(a[] string)[]string{

	var temp[]string
	for i:=len(a)-1;i>=0;i-- {
		temp=append(temp, a[i])
	}
	return temp
}

func rev(a[] int)[]int{

	var temp[]int
	for i:=len(a)-1;i>=0;i-- {
		temp=append(temp, a[i])
	}
	return temp
}

func main () {

	array1:=[]int{1,2,3,4,5,5,78,0}
	fmt.Println("array of numbers:",array1)
	slices.Reverse(array1)
	fmt.Println(array1)

	array2:=[]string{"chintu","vijay","chintuVijay"}
	fmt.Println("array of string :",array2)
	slices.Reverse(array2)
	fmt.Println(array2)

	array3:=[]string{"chintu","vijay","chintu-Vijay"}
	fmt.Println("array of string :",array3)
	output:=revString(array3)
	fmt.Println(output)

	array4:=[]int{1,2,3,4,5,5,78,0}
	fmt.Println("array of numbers:",array4)
	output1:=rev(array4)
	fmt.Println(output1)

}