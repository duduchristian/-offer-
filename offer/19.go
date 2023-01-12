package offer

//面试题19：最多删除一个字符得到回文
//题目：给定一个字符串，请判断如果最多从字符串中删除一个字符能不能得到一个回文字符串。
//例如，如果输入字符串"abca"，由于删除字符'b'或'c'就能得到一个回文字符串，因此输出为true。

func ValidPalindrome(s string) bool {
	start, stop := 0, len(s)-1
	for start < stop {
		if s[start] != s[stop] {
			break
		}
		start++
		stop--
	}
	return isPalindrome(s, start+1, stop) ||
		isPalindrome(s, start, stop-1)
}

func isPalindrome(s string, start, stop int) bool {
	for start < stop {
		if s[start] != s[stop] {
			break
		}
		start++
		stop--
	}
	return start >= stop
}
