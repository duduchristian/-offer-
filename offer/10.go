package offer

//面试题10：和为k的子数组
//题目：输入一个整数数组和一个整数k，请问数组中有多少个数字之和等于k的连续子数组？
//例如，输入数组[1，1，1]，k的值为2，有2个连续子数组之和等于2。

func SubarraySum(nums []int, k int) int {
	res := 0
	mem := make(map[int]int, len(nums)+1)
	mem[0] = 1
	curSum := 0
	for _, num := range nums {
		curSum += num
		count := mem[curSum-k]
		mem[curSum]++
		res += count
	}
	return res
}
