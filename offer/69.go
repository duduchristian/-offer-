package offer

//面试题69：山峰数组的顶部

func PeakIndexInMountainArray(nums []int) int {
	left, right := 1, len(nums)-2
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}
		if nums[mid] > nums[mid-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
