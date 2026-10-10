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
}