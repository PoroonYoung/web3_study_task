package main

import (
	"fmt"
)

//
//第一反应是转换成字符串再转数字，加一后转回来。
//后来发现这个考察的不是转换，当数字超过int256范围之后转换数字就会变成0
//想要的其实是手动模拟十进制的加法进位
//func plusOne(digits []int) []int {
//	str := strconv.Itoa(digits[0])
//	for i := 1; i < len(digits); i++ {
//		str = str + strconv.Itoa(digits[i])
//	}
//	num, _ := strconv.ParseInt(str, 10, 256)
//	num++
//	strAfterPlus := strconv.FormatInt(num, 10)
//	var array []int
//	for _, v := range strAfterPlus {
//		array = append(array, int(v-'0'))
//	}
//	return array
//}
//

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			break
		}
	}
	allZero := true
	for i := range digits {
		if digits[i] != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		digits = make([]int, len(digits)+1)
		digits[0] = 1
	}
	return digits
}
func main() {
	fmt.Println(plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
	fmt.Println(plusOne([]int{9}))
}
