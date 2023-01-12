package offer

//试题12：左右两边子数组的和相等
//题目：输入一个整数数组，如果一个数字左边的子数组的数字之和等于右边的子数组的数字之和，那么返回该数字的下标。
//如果存在多个这样的数字， 则返回最左边一个数字的下标。如果不存在这样的数字，则返回-1。例如，在数组[1，7，3，6，2，9]中，
//下标为3的数字（值为6）的左边3个数字1、7、3的和与右边两个数字2和9的和相等，都是11，因此正确的输出值是3。

func PivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for i, num := range nums {
		sum += num
		if sum-num == total-sum {
			return i
		}
	}
	return -1
}
