package offer

import (
	"algorithm/utils"
	"github.com/emirpasic/gods/queues/priorityqueue"
)

//面试题61：和最小的k个数对
//题目：给定两个递增排序的整数数组，从两个数组中各取一个数字u和v组成一个数对（u，v），请找出和最小的k个数对。
//例如，输入两个数组[1，5，13，21]和[2，4，9，15]，和最小的3个数对为（1，2）、（1，4）和（2，5）。

func KSmallestPairs(nums1, nums2 []int, k int) [][2]int {
	minHeap := priorityqueue.NewWith(func(a, b interface{}) int {
		p1 := a.([2]int)
		p2 := b.([2]int)
		return nums1[p1[0]] + nums2[p1[1]] - nums1[p2[0]] - nums2[p2[1]]
	})
	if len(nums2) > 0 {
		for i := 0; i < utils.Min(k, len(nums1)); i++ {
			minHeap.Enqueue([2]int{i, 0})
		}
	}

	var res [][2]int
	for ; k > 0 && !minHeap.Empty(); k-- {
		p, _ := minHeap.Dequeue()
		indices := p.([2]int)
		res = append(res, [2]int{nums1[indices[0]], nums2[indices[1]]})

		if indices[1] < len(nums2)-1 {
			minHeap.Enqueue([2]int{indices[0], indices[1] + 1})
		}
	}
	return res
}
