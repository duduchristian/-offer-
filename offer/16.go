package offer

import (
	"algorithm/utils"
)

//面试题16：不含重复字符的最长子字符串
//题目：输入一个字符串，求该字符串中不含重复字符的最长子字符串的长度。
//例如，输入字符串"babcca"，其最长的不含重复字符的子字符串是"abc"，长度为3。

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var counts [256]int
	longest := 1
	left := 0
	isDup := false

	for right := 0; right < len(s); right++ {
		rightByte := s[right]
		counts[rightByte]++
		if counts[rightByte] == 2 {
			isDup = true
		}
		for isDup {
			leftByte := s[left]
			counts[leftByte]--
			if leftByte == rightByte {
				isDup = false
			}
			left++
		}

		longest = utils.Max(longest, right-left+1)
	}

	return longest
}
