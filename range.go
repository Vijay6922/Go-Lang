package main

import "fmt"

func main() {

	//range uses

	for k, v := range "vijay" {
		// fmt.Println(k,v)
		fmt.Println(k,string(v))
	}

	m := make(map[string]string)
	m["name"] = "vijay"
	m["second name"] = "kumar"
	fmt.Println(m)
	for k,v := range m {
		fmt.Println(k,v)
	}
}