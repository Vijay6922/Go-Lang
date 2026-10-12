package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	url := "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close() //prevent memory leaks

	if response.StatusCode == http.StatusOK {
		// bodyBytes, err := io.ReadAll(response.Body) //reading response body
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// data := string(bodyBytes)
		// fmt.Println(data)

		// todoItem := todo{}
		// json.Unmarshal(bodyBytes, &todoItem) //converting the bodybytes into todo

		// fmt.Printf("data from API : %+v", todoItem)

		// todoItem := todo{}
		// json.NewDecoder(response.Body).Decode(&todoItem) //converting the bodybytes into todo
		// fmt.Printf("data from API : %+v", todoItem)

		todoItem := todo{}

		decoder := json.NewDecoder(response.Body)
		decoder.DisallowUnknownFields() //if any feilds are missing gives an error
		if err := decoder.Decode(&todoItem); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("data from API : %+v", todoItem)

	}

}
