package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task of id :",id)
}

func main() {

	for i:=0;i<=10;i++ {
		//run the task parellelly 
		go task(i)

		//anonymous goroutines 
		go func (i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second* 2)
}
