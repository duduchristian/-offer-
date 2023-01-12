package offer

import "container/list"

type Stack struct {
	*list.List
}

func NewStack() Stack {
	return Stack{
		list.New(),
	}
}

func (s Stack) Push(value interface{}) {
	s.PushBack(value)
}

func (s Stack) Pop() interface{} {
	e := s.Back()
	s.Remove(e)
	return e.Value
}

func (s Stack) Empty() bool {
	return s.Len() == 0
}

func (s Stack) Peek() interface{} {
	return s.Back().Value
}

func (s Stack) PushInt(v int) {
	s.Push(v)
}

func (s Stack) PopInt() int {
	return s.Pop().(int)
}

func (s Stack) PeekInt() int {
	return s.Peek().(int)
}
