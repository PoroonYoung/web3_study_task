package main

import "fmt"

func removeDuplicates(nums []int) int {
	m := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			nums = nums[:len(nums)-1]
			i--
		} else {
			m[nums[i]] = true
		}
	}
	fmt.Println(nums)
	return len(nums)
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 2, 3}))
}
