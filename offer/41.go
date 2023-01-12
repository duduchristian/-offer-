package offer

//面试题41：滑动窗口的平均值

type MovingAverage struct {
	nums     Queue
	capacity int
	sum      int
}

func NewMovingAverage(size int) *MovingAverage {
	return &MovingAverage{
		nums:     NewQueue(),
		capacity: size,
	}
}

func (p *MovingAverage) Next(val int) float64 {
	p.nums.OfferInt(val)
	p.sum += val
	if p.nums.Len() > p.capacity {
		p.sum -= p.nums.PollInt()
	}
	return float64(p.sum) / float64(p.nums.Len())
}
