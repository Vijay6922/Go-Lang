package main

import "fmt"

func printNumbers(num int) {
    if num <= 100 {
        fmt.Print(num, "\t")
        printNumbers(num + 1)
    }
}
func main() {

    number := 1
    printNumbers(number)
}