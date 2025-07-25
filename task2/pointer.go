package main

import "fmt"

// 传入int指针，让指向数字+1
func plus1(intPointer *int) {
	*intPointer += 1
}

// 传入int切片的指针，每个int都*2
func multiply2(slicePointer *[]int) {
	ints := *slicePointer
	for i, _ := range ints {
		ints[i] *= 2
	}
}

func main() {
	num := 1
	plus1(&num)
	fmt.Println(num)
	nums := []int{1, 2, 3}
	multiply2(&nums)
	fmt.Println(nums)
}
