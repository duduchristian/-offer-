package offer

//面试题70：排序数组中只出现一次的数字

func SingleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)/2
	for left <= right {
		mid := (left + right) / 2
		i := mid * 2
		if i < len(nums)-1 && nums[i] != nums[i+1] {
			if mid == 0 || nums[i-2] == nums[i-1] {
				return nums[i]
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[len(nums)-1]
}
