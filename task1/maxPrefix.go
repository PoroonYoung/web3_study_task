package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {

	minLen := 0
	minStr := ""
	for _, str := range strs {
		if minLen == 0 {
			minLen = len(str)
			minStr = str
			continue
		}
		if len(str) < minLen {
			minStr = str
			minLen = len(str)
		}
	}
	return getPrefix(strs, minStr)
}

func getPrefix(strs []string, minStr string) string {
	if len(strs) == 0 || minStr == "" {
		return ""
	}
	for _, str := range strs {
		if !strings.HasPrefix(str, minStr) {
			minStr = minStr[:len(minStr)-1]
			minStr = getPrefix(strs, minStr)
		}
	}
	return minStr
}

func main() {
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs1))
	strs2 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs2))
}
