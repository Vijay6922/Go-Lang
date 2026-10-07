package main

import "fmt"

func calculator(a float64, b float64, action string) float64 {

	switch action {
	case "add":
		{
			return a + b
		}
	case "sub":
		{
			if a > b {
				fmt.Println("you may get the negative value")
				return a-b
			} else {
			return a-b
		}
	}
	case "mul" : {
		return a*b
	}
	case "div" : {
		return a/b
	}
	}
	return 0
}

//multiple return types return 
func language() (string,string,string) {
	return "telugu","english","hindi"
}

//pass function as param
func process(fn func(a int) int) {
	fn(1)
}

// return function as return type
func process2() func(a int) int {
	return func(a int) int {
		return 1
	}
}

func main() {

	cal := calculator(3,2,"div")
	fmt.Println(cal)
	lang,lang2,lang3 := language()
	fmt.Println(lang,lang2,lang3)
	_,aa,bb := language()
	fmt.Println(aa,bb)

	//passing fucntion
	fn:= func(a int)int {
		return 1
	}
	process(fn)

	//returning function 
	fn1 := process2()
	fn1(4)


}
