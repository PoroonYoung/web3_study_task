package main

import (
	"fmt"
	"sync"
)

var count int

func countWithLock(lock *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	count++
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	lock := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go countWithLock(&lock, &wg)
	}
	wg.Wait()
	fmt.Println("count=", count)
}
