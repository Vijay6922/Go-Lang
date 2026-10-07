package main
import "fmt"

func main(){

	// noOfPrime :=
	 var i = prime(200,300)
	 fmt.Println("no of prime numbers : ",i)

}

func prime( startingIndex int , endIndex int) int{

	primes := 0
	for i:=startingIndex; i<= endIndex;i++ {
		counter := 0
		for j :=1; j<=i; j++ {
			if(i%j==0){
				counter = counter+1
				// fmt.Println(i,j,startingIndex)
			}
		}
		if counter <= 2 {
			fmt.Println(i)
			primes++
		}
	}
	return primes
	// return counter
}
