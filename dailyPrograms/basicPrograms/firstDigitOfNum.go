package main 
import "fmt"
func main () {

	num := 2345

	temp := 0
	rev:=0

	for {
	if num != 0 {
		rem := num % 10
		rev = (rev*10)+rem
		num = num / 10
	} else {
		break
	}
	}
	temp = rev
	r:= temp% 10
	fmt.Println("First digit of the number is :",r)

	// num := 2000 // Test with a number ending in zeros
	// This loop strips away digits from the right side
	// for {
	// 	if num >= 10 { // Bug Fix: Stop when only 1 digit is left
	// 		num = num / 10
	// 	} else {
	// 		break
	// 	}
	// }
	// 'num' now naturally holds the first digit directly
	// fmt.Println("First digit of the number is :", num)
}