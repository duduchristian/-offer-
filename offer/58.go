package offer

import "github.com/emirpasic/gods/maps/treemap"

//面试题58：日程表
//题目：请实现一个类型MyCalendar用来记录自己的日程安排，该类型用方法book（int start，int end）在日程表中添加一个时间区域为[start，end）的事项
//（这是一个半开半闭区间）。如果[start，end）中之前没有安排其他事项，则成功添加该事项并返回true；否则，不能添加该事项，并返回false。

type MyCalendar struct {
	m *treemap.Map
}

func NewMyCalendar() *MyCalendar {
	return &MyCalendar{treemap.NewWithIntComparator()}
}

func (p *MyCalendar) Book(start, end int) bool {
	_, beforeEnd := p.m.Floor(start)
	afterStart, _ := p.m.Ceiling(start)
	if beforeEnd != nil && beforeEnd.(int) > start {
		return false
	}
	if afterStart != nil && afterStart.(int) < end {
		return false
	}

	p.m.Put(start, end)
	return true
}
