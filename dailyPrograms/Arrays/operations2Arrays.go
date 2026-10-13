package main

import (
	"fmt"
)

func addition(a []int, b []int) []int {

	sum := make([]int,len(a))
	if len(a) != len(b) {
		fmt.Println("not possible to add as length is difference")
		return nil
	} else {
	for i := 0; i < len(a); i++ {
		sum[i]=a[i]+b[i]
	}
	return sum
	}
}

func sub(a []int, b []int) []int {

	sum := make([]int,len(a))
	if len(a) != len(b) {
		fmt.Println("not possible to add as length is difference")
		return nil
	} else {
	for i := 0; i < len(a); i++ {
		sum[i]=a[i]-b[i]
	}
	return sum
	}
}

func mul(a []int, b []int) []int {

	sum := make([]int,len(a))
	if len(a) != len(b) {
		fmt.Println("not possible to add as length is difference")
		return nil
	} else {
	for i := 0; i < len(a); i++ {
		sum[i]=a[i]*b[i]
	}
	return sum
	}
}

func div(a []int, b []int) []int {

	sum := make([]int,len(a))
	if len(a) != len(b) {
		fmt.Println("not possible to add as length is difference")
		return nil
	} else {
	for i := 0; i < len(a); i++ {
		sum[i]=a[i]/b[i]
	}
	return sum
	}
}

func mod(a []int, b []int) []int {

	sum := make([]int,len(a))
	if len(a) != len(b) {
		fmt.Println("not possible to add as length is difference")
		return nil
	} else {
	for i := 0; i < len(a); i++ {
		sum[i]=a[i]%b[i]
	}
	return sum
	}
}

func avg(a []int) int {

	avg :=0
	for i := 0; i < len(a); i++ {
		avg+=a[i]
	}
	return avg/len(a)
}


func main() {

	array1 := []int{12, 2, 344, 4, 5}
	array2 := []int{9, 9, 9, 9, 9}
	operation := "avg"   //can take from the user 
	switch (operation){
	case "Addtion" : {
		fmt.Println("addition of 2 arrays :",addition(array1,array2))
	}
	case "sub" : {
		fmt.Println("subtraction of 2 arrays :",sub(array1,array2))
	}
	case "mul" : {
		fmt.Println("multiplication of 2 arrays :",mul(array1,array2))
	}
	case "div" : {
		fmt.Println("division of 2 arrays :",div(array1,array2))
	}
	case "mod" : {
		fmt.Println("modulus of 2 arrays :",mod(array1,array2))
	}
	case "avg" : {
		fmt.Println("avg of array :",avg(array1))
	}
	default :{
		fmt.Println("addition of 2 arrays :",addition(array1,array2))
	}
	}

}