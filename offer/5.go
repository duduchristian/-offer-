package offer

//面试题5：单词长度的最大乘积
//题目：输入一个字符串数组words，请计算不包含相同字符的两个字符串words[i]和words[j]的长度乘积的最大值。
//如果所有字符串都包含至少一个相同字符，那么返回0。假设字符串中只包含英文小写字母。
//例如，输入的字符串数组words为["abcw"，"foo"，"bar"，"fxyz"，"abcdef"]，数组中的字符串"bar"与"foo"没有相同的字符，
//它们长度的乘积为9。"abcw"与"fxyz"也没有相同的字符，它们长度的乘积为16，这是该数组不包含相同字符的一对字符串的长度乘积的最大值。

func MaxProduct(words []string) (res int) {
	mem := make([]int32, len(words))
	for i, word := range words {
		for _, c := range word {
			mem[i] |= 1 << (c - 'a')
		}
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if mem[i]&mem[j] == 0 {
				max := len(words[i]) * len(words[j])
				if max > res {
					res = max
				}
			}
		}
	}

	return
}
