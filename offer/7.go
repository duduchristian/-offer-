package offer

import "sort"

//面试题7：数组中和为0的3个数字
//题目：输入一个数组，如何找出数组中所有和为0的3个数字的三元组？需要注意的是，返回值中不得包含重复的三元组。
//例如，在数组[-1，0，1，2，-1，-4]中有两个三元组的和为0，它们分别是[-1，0，1]和[-1，-1，2]。

func ThreeSum(numbers []int) (results [][3]int) {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	for i := 0; i < len(numbers)-2; {
		twoSum(numbers, i, &results)
		temp := numbers[i]
		for temp == numbers[i] && i < len(numbers)-2 {
			i++
		}
	}
	return results
}

func twoSum(nums []int, i int, results *[][3]int) {
	j, k := i+1, len(nums)-1
	for j < k {
		sum := nums[i] + nums[j] + nums[k]
		if sum == 0 {
			*results = append(*results, [3]int{nums[i], nums[j], nums[k]})
			temp := nums[j]
			for temp == nums[j] && j < k {
				j++
			}
		} else if sum < 0 {
			j++
		} else {
			k--
		}
	}
}
