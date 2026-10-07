package main

import (
	"fmt"
)

func sendData(message chan int) {
	fmt.Println("sending data : ", <-message)
}

func recieveData(message chan int, num1 int, num2 int) {
	result := num1* num2
	message <- result
}

func main() {

	// messageChan := make(chan string,1)

	// //data sent into channel 
	// messageChan <- "vijay"
	// fmt.Println(messageChan)

	// //recieving data form channel 
	// msg := <- messageChan
	// fmt.Println(msg)

	//sending data to channel function
	messagesend := make(chan int)
	go sendData(messagesend)
	messagesend <- 10
	// time.Sleep(time.Second*2)  // need for sending data not for recieving 

	//receive the data from function
	recieveMessage := make(chan int)
	go recieveData(recieveMessage,5,9)
	result := <-recieveMessage
	fmt.Println("recieving data from channel :",result)

}
