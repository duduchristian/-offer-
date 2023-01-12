package offer

//面试题32：有效的变位词
//题目：给定两个字符串s和t，请判断它们是不是一组变位词。在一组变位词中，它们中的字符及每个字符出现的次数都相同， 但字符的顺序不能相同。
//例如，"anagram"和"nagaram"就是一组变位词。

func IsAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	counts := map[rune]int{}
	for _, r := range str1 {
		counts[r] += 1
	}

	for _, r := range str2 {
		if count := counts[r]; count == 0 {
			return false
		}
		counts[r]--
	}
	return true
}
