package main

import "fmt"

func main() {
	intArray := []int{1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6}
	countMap := make(map[int]int)
	for _, item := range intArray {
		countMap[item]++
	}
	for _, item := range intArray {
		if countMap[item] == 1 {
			fmt.Println(item)
		}
	}
}
