package offer

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

//面试题59：数据流的第k大数字

type KthLargest struct {
	pq   *priorityqueue.Queue
	size int
}

func NewKthLargest(k int, nums []int) *KthLargest {
	ret := &KthLargest{
		pq:   priorityqueue.NewWith(utils.IntComparator),
		size: k,
	}
	for _, num := range nums {
		ret.Add(num)
	}
	return ret
}

func (p *KthLargest) Add(val int) int {
	p.pq.Enqueue(val)
	for p.pq.Size() > p.size {
		p.pq.Dequeue()
	}
	ret, _ := p.pq.Peek()
	return ret.(int)
}
