package main

import "fmt"
func main () {

	num := 50 

	for i:=1;i<=num;i++ {
		if num%i==0{
			fmt.Println(i)
		}
	}
}