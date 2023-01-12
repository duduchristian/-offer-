package offer

import (
	"algorithm/utils"
	"sort"
)

//面试题34：外星语言是否排序
//题目：有一门外星语言，它的字母表刚好包含所有的英文小写字母，只是字母表的顺序不同。
//给定一组单词和字母表顺序，请判断这些单词是否按照字母表的顺序排序。
//例如，输入一组单词["offer"，"is"，"coming"]，以及字母表顺序"zyxwvutsrqponmlkjihgfedcba"，
//由于字母'o'在字母表中位于'i'的前面，因此单词"offer"排在"is"的前面；
//同样，由于字母'i'在字母表中位于'c'的前面，因此单词"is"排在"coming"的前面。
//因此，这一组单词是按照字母表顺序排序的，应该输出true。

func IsAlienSorted(words []string, order string) bool {
	orderArray := make([]int, 26)
	for i := range order {
		orderArray[order[i]-'a'] = i
	}
	return sort.SliceIsSorted(words, func(i, j int) bool {
		minLength := utils.Min(len(words[i]), len(words[j]))
		for index := 0; index < minLength; index++ {
			c1 := words[i][index]
			c2 := words[j][index]
			o1 := orderArray[c1-'a']
			o2 := orderArray[c2-'a']
			if o1 == o2 {
				continue
			}
			if o1 < o2 {
				return true
			}
			return false
		}
		return len(words[i]) < len(words[j])
	})
}
