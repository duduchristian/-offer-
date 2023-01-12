package offer

//面试题14：字符串中的变位词
//题目：输入字符串s1和s2，如何判断字符串s2中是否包含字符串s1的某个变位词？如果字符串s2中包含字符串s1的某个变位词，
//则字符串s1至少有一个变位词是字符串s2的子字符串。假设两个字符串中只包含英文小写字母。例如，字符串s1为"ac"，
//字符串s2为"dgcaf"，由于字符串s2中包含字符串s1的变位词"ca"，因此输出为true。如果字符串s1为"ab"，字符串s2为"dgcaf"，则输出为false。

func CheckInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var counts [26]int
	for i := 0; i < len(s1); i++ {
		counts[s1[i]-'a']++
		counts[s2[i]-'a']--
	}
	if areAllZeros(&counts) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		counts[s2[i]-'a']--
		counts[s2[i-len(s1)]-'a']++
		if areAllZeros(&counts) {
			return true
		}
	}
	return false
}

func areAllZeros(counts *[26]int) bool {
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}
