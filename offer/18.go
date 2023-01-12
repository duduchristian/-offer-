package offer

import "algorithm/utils"

//面试题18：有效的回文
//题目：给定一个字符串，请判断它是不是回文。假设只需要考虑字母和数字字符，并忽略大小写。
//例如，"Was it a cat I saw？"是一个回文字符串，而"race a car"不是回文字符串。

func IsPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		ch1, ch2 := s[i], s[j]
		if !utils.IsLetterOrDigit(ch1) {
			i++
		} else if !utils.IsLetterOrDigit(ch2) {
			j--
		} else {
			ch1 = utils.ToLowerCase(ch1)
			ch2 = utils.ToLowerCase(ch2)
			if ch1 != ch2 {
				return false
			}
			i++
			j--
		}
	}
	return true
}
