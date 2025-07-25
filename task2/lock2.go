package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var num int32

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&num, 1)
		}()
	}
	wg.Wait()
	fmt.Println(num)
}
