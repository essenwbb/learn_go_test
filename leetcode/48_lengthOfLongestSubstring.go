package leetcode

import "math"

func lengthOfLongestSubstringCrazy(s string) int {
	sRune := []rune(s)
	length := len(sRune)
	if length <= 1 {
		return length
	}
	max := 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j <= length; j++ {
			tmp := make(map[rune]bool, j-i)
			for offset := 0; i+offset < j; offset++ {
				tmp[rune(sRune[i+offset])] = true
			}
			if len(tmp) == j-i {
				max = int(math.Max(float64(max), float64(j-i)))
			}
		}
	}
	return max
}

func lengthOfLongestSubstring(s string) int {
	return 0
}
