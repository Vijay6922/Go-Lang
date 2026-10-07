package main

import "fmt"

func main() {

	//two channels
	chan1 := make(chan int)
	chan2 := make(chan string)

	//using closure  not passing any value in function 
	go func() {
		chan1 <- 12
	}()

	go func() {
		chan2 <- "vijay"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Chan1 value",chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Chan2 value",chan2Val)
		}
	}
}
