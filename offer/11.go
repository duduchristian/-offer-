package offer

import "algorithm/utils"

//面试题11：0和1个数相同的子数组
//题目：输入一个只包含0和1的数组，请问如何求0和1的个数相同的最长连续子数组的长度？
//例如，在数组[0，1，0]中有两个子数组包含相同个数的0和1，分别是[0，1]和[1，0]，它们的长度都是2，因此输出2。

func FindMaxLength(nums []int) int {
	res := 0
	curSum := 0
	mem := make(map[int]int, len(nums)+1)
	mem[0] = -1
	for i, num := range nums {
		if num == 0 {
			curSum--
		} else {
			curSum++
		}
		if length, ok := mem[curSum]; ok {
			res = utils.Max(res, i-length)
		} else {
			mem[curSum] = i
		}
	}
	return res
}
