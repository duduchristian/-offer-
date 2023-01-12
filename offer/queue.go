package offer

import "container/list"

type Queue struct {
	*list.List
}

func NewQueue() Queue {
	return Queue{
		list.New(),
	}
}

func (q Queue) Offer(val interface{}) {
	q.PushFront(val)
}

func (q Queue) Poll() interface{} {
	e := q.Back()
	q.Remove(e)
	return e.Value
}

func (q Queue) Peek() interface{} {
	return q.Back().Value
}

func (q Queue) Empty() bool {
	return q.Len() == 0
}

func (q Queue) OfferInt(val int) {
	q.Offer(val)
}

func (q Queue) PollInt() int {
	return q.Poll().(int)
}

func (q Queue) PeekInt() int {
	return q.Peek().(int)
}
