package offer

import (
	"algorithm/utils"
	"math"
)

//面试题8：和大于或等于k的最短子数组
//题目：输入一个正整数组成的数组和一个正整数k，请问数组中和大于或等于k的连续子数组的最短长度是多少？如果不存在所有数字之和大于或等于k的子数组，则返回0。
//例如，输入数组[5，1，4，3]，k的值为7，和大于或等于7的最短连续子数组是[4，3]，因此输出它的长度2。

func MinSubArrayLen(k int, nums []int) int {
	left, curSum := 0, 0
	res := math.MaxInt
	for right := 0; right < len(nums); right++ {
		curSum += nums[right]
		for curSum >= k {
			res = utils.Min(res, right-left+1)
			curSum -= nums[left]
			left++
		}
	}
	if res == math.MaxInt {
		res = 0
	}
	return res
}

func MinSubArrayLenBruteForce(k int, nums []int) int {
	res := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= k {
				res = utils.Min(j-i+1, res)
				break
			}
		}
	}
	if res == math.MaxInt {
		res = 0
	}
	return res
}
