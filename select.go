package main

import (
	"fmt"
	"time"
)

func main() {

	chan1 := make(chan int)
	chan2 := make(chan int)

	go func () {
	for i := range 3 {
		chan1 <- i
		time.Sleep(time.Second)
	}
	}()

	go func () {
	for i := range 3 {
		chan2 <- i
		time.Sleep(time.Second)
	}
	}()

	for range 6{
	select {
	case value := <- chan1 :
		fmt.Println("channel 1", value)
	case value := <- chan2 : 
		fmt.Println("channel 2", value)
	}
	}
}
