package main

import (
	"fmt"
	"math"
)

func min(strs []string) int {
	min := math.MaxInt
	for _, s := range strs {
		if len(s) < min {min = len(s)}
	}
	return min
}

func isEqIdx(strs []string, idx int) bool {
	for i:=0; i<len(strs); i++ {
		sym := strs[0][idx]
		if !(sym == strs[i][idx]) {
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	max_symbols := min(strs)
	count := 0
	for i:=0; i<max_symbols; i++ {
		if isEqIdx(strs, i) {
			count++
		} else {
			break
		}
	}
	return strs[0][0:count]
}

func main() {
	test_cases := [][]string {
		{"a"},
		{"flower","flow","flight"},
		{"dog","racecar","car"},
	}
	for _, test_case := range test_cases {
		result := longestCommonPrefix(test_case)
		fmt.Println(result)
	}
}
