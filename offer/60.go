package offer

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
)

//面试题60：出现频率最高的k个数字
//题目：请找出数组中出现频率最高的k个数字。
//例如，当k等于2时，输入数组[1，2，2，1，3，1]，由于数字1出现了3次，数字2出现了2次，数字3出现了1次，因此出现频率最高的2个数字是1和2。

func TopKFrequent(nums []int, k int) []int {
	numToCount := map[int]int{}
	for _, num := range nums {
		numToCount[num]++
	}
	minHeap := priorityqueue.NewWith(func(a, b interface{}) int {
		val1 := a.(int)
		val2 := b.(int)
		if numToCount[val1] < numToCount[val2] {
			return -1
		}
		if numToCount[val1] > numToCount[val2] {
			return 1
		}
		return val2 - val1
	})
	for num := range numToCount {
		minHeap.Enqueue(num)
		for minHeap.Size() > k {
			minHeap.Dequeue()
		}
	}
	ret := make([]int, 0, k)
	for minHeap.Size() > 0 {
		num, _ := minHeap.Dequeue()
		ret = append(ret, num.(int))
	}
	return ret
}
