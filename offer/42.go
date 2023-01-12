package offer

//面试题42：最近请求次数

type RecentAverage interface {
	Ping(t int) int
}

type RecentAverageImpl struct {
	times      Queue
	windowSize int
}

func NewRecentAverage() RecentAverage {
	return &RecentAverageImpl{
		times:      NewQueue(),
		windowSize: 3000,
	}
}

func (p *RecentAverageImpl) Ping(t int) int {
	p.times.OfferInt(t)
	for p.times.PeekInt()+p.windowSize < t {
		p.times.Poll()
	}
	return p.times.Len()
}
