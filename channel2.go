package main

import (
	"fmt"
	"time"
)

// Signals completion to caller via channel
func donefunc(message chan bool) {
	defer func() { message <- true }() // Signal done on function exit
	fmt.Print("processing")
	fmt.Println()
}

// Worker: processes emails until channel is closed
func sendemail(message chan string, comp chan bool) {
	defer func() { comp <- true }() // Signal main when worker finishes
	for mail := range message {     // Reads until channel closes
		fmt.Println("sending email to ", mail)
		time.Sleep(time.Second)
	}
}

func main() {
	// 1. Basic Goroutine synchronization
	done := make(chan bool)
	go donefunc(done)
	<-done // Block until donefunc completes

	// 2. Producer-Consumer pattern
	email := make(chan string) // Data channel
	comp := make(chan bool)    // Completion signal channel

	go sendemail(email, comp)  // Start worker

	// Send emails to worker
	for i := 0; i <= 10; i++ {
		email <- fmt.Sprintf("%d@gmail.com", i)
	}
	// email <-"vijay@gmail.com"
	fmt.Println("done sending")
	close(email) // Signals worker to exit loop

	<-comp // Wait for worker to finish
}
