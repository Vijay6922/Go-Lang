package main

import (
	"fmt"
	"sync"
)

type example struct {
	id int
	mu sync.Mutex
}

func (ex *example) sum(wg *sync.WaitGroup) {
	defer func(){
		ex.mu.Unlock()
		wg.Done()
	}()
	ex.mu.Lock()
	ex.id += 1
}

func main() {

	var wg sync.WaitGroup
	ex1 := example{id: 0}

	for i:=0;i<100;i++ {
		wg.Add(1)
		go ex1.sum(&wg)
	}
	wg.Wait()
	fmt.Println(ex1.id)


}
