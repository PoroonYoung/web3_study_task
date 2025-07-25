package main

import (
	"fmt"
	"sync"
)

func main() {

	intChan := make(chan int, 10)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(chan int) {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			intChan <- i
		}
		close(intChan)
	}(intChan)

	go func(chan int) {
		defer wg.Done()
		for i := range intChan {
			fmt.Println(i)
		}
	}(intChan)

	wg.Wait()
}
