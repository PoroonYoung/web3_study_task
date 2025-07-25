package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	intChan := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	go func() {
		defer wg.Done()
		for v := range intChan {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
