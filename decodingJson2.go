package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type todos struct {
	UserID   int    `json:"-"`
	ID        int    `json:"-"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed"`
}

func main() {

		todoItems := &todos{1,1,"",false}

		//convert struct to JSON
		todo,err:= json.Marshal(todoItems)
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Println(string(todo))
	}
