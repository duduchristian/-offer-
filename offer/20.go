package offer

//面试题20：回文子字符串的个数
//题目：给定一个字符串，请问该字符串中有多少个回文连续子字符串？
//例如，字符串"abc"有3个回文子字符串，分别为"a"、"b"和"c"；而字符串"aaa"有6个回文子字符串，分别为"a"、"a"、"a"、"aa"、"aa"和"aaa"。

func CountSubstrings(s string) int {
	ret := 0
	for i := 0; i < len(s); i++ {
		ret += countPalindrome(s, i, i)
		ret += countPalindrome(s, i, i+1)
	}
	return ret
}

func countPalindrome(s string, start, stop int) int {
	count := 0
	for start >= 0 && stop < len(s) && s[start] == s[stop] {
		count++
		start--
		stop++
	}
	return count
}
