package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]string)
	

	m["name"] = "vijay"
	m["second name"] = "kumar"
	fmt.Println(m)
	m1 := map[string]string {}
	maps.Copy(m1,m)  //m1 is the new map and m is the old one 
	

	fmt.Println(len(m1))
	m1["age"] = "24"
	fmt.Println(maps.Equal(m1,m))
	fmt.Println(m1)
}