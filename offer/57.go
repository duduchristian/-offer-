package offer

import (
	"github.com/emirpasic/gods/maps/treemap"
)

//面试题57：值和下标之差都在给定的范围内
//题目：给定一个整数数组nums和两个正数k、t，请判断是否存在两个不同的下标i和j满足i和j之差的绝对值不大于给定的k，
//并且两个数值nums[i]和nums[j]的差的绝对值不大于给定的t。

//例如，如果输入数组{1，2，3，1}，k为3，t为0，由于下标0和下标3对应的数字之差的绝对值为0，因此返回true。
//如果输入数组{1，5，9，1，5，9}，k为2，t为3，由于不存在两个下标之差小于或等于2且它们差的绝对值小于或等于3的数字，因此此时应该返回false。

func ContainsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	buckets := map[int]int{}
	bucketSize := t + 1
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		id := getBucketID(num, bucketSize)

		if _, ok := buckets[id]; ok {
			return true
		}

		if _, ok := buckets[id-1]; ok && num-buckets[id-1] <= t {
			return true
		}

		if _, ok := buckets[id+1]; ok && buckets[id+1]-num <= t {
			return true
		}

		buckets[id] = num
		if i >= k {
			delete(buckets, getBucketID(nums[i-k], bucketSize))
		}
	}
	return false
}

func getBucketID(num, bucketSize int) int {
	if num >= 0 {
		return num / bucketSize
	}
	return (num+1)/bucketSize - 1
}

func ContainsNearbyAlmostDuplicateB(nums []int, k, t int) bool {
	m := treemap.NewWithIntComparator()
	for i := 0; i < len(nums); i++ {
		lower, _ := m.Floor(nums[i])
		if lower != nil && nums[i]-lower.(int) <= t {
			return true
		}
		upper, _ := m.Ceiling(nums[i])
		if upper != nil && upper.(int)-nums[i] <= t {
			return true
		}
		if i >= k {
			m.Remove(nums[i-k])
		}
		m.Put(nums[i], struct{}{})
	}
	return false
}
