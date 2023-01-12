package offer

import "sort"

//面试题33：变位词组
//题目：给定一组单词，请将它们按照变位词分组。例如，输入一组单词["eat"，"tea"，"tan"，"ate"，"nat"，"bat"]，
//这组单词可以分成3组，分别是["eat"，"tea"，"ate"]、["tan"，"nat"]和["bat"]。假设单词中只包含英文小写字母。

func GroupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, str := range strs {
		arr := []byte(str)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		sorted := string(arr)
		groups[sorted] = append(groups[sorted], str)
	}
	ret := make([][]string, 0, len(groups))
	for _, group := range groups {
		ret = append(ret, group)
	}
	return ret
}
