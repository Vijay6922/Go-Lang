package main 


import "fmt" 

//perfect number is a positive integer that is equal to the sum of its proper positive divisors, excluding the number itself. 
// For example, 6 is a perfect number because its divisors are 1, 2, and 3, and 1 + 2 + 3 = 6.
func main () {

	num := 28
	temp := num
	sum:=0

	for i:=1;i<num;i++ {

		if num%i==0 {
			sum += i
		}
	}
	if sum == temp {
		fmt.Println(num, "is a perfect number")
	} else {
		fmt.Println(num, "is not a perfect number")
	}

}