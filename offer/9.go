package offer

//面试题9：乘积小于k的子数组
//题目：输入一个由正整数组成的数组和一个正整数k，请问数组中有多少个数字乘积小于k的连续子数组？
//例如，输入数组[10，5，2，6]，k的值为100，有8个子数组的所有数字的乘积小于100，它们分别是[10]、[5]、[2]、[6]、[10，5]、[5，2]、[2，6]和[5，2，6]。

func NumSubarrayProductLessThanK(nums []int, k int) int {
	left, curProduct := 0, 1
	res := 0
	for right := 0; right < len(nums); right++ {
		curProduct *= nums[right]
		for curProduct >= k {
			curProduct /= nums[left]
			left++
		}
		res += right - left + 1
	}
	return res
}
