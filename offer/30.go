package offer

import (
	"math/rand"
	"time"
)

//面试题30：插入、删除和随机访问都是O（1）的容器
//题目：设计一个数据结构，使如下3个操作的时间复杂度都是O（1）。

type RandomizedSet struct {
	numToLocation map[int]int
	nums          []int
}

func NewRandomizedSet() *RandomizedSet {
	return &RandomizedSet{
		numToLocation: map[int]int{},
		nums:          nil,
	}
}

func (p *RandomizedSet) Insert(val int) bool {
	if _, ok := p.numToLocation[val]; ok {
		return false
	}
	p.numToLocation[val] = len(p.nums)
	p.nums = append(p.nums, val)
	return true
}

func (p *RandomizedSet) Remove(val int) bool {
	if location, ok := p.numToLocation[val]; ok {
		return false
	} else {
		p.numToLocation[p.nums[len(p.nums)-1]] = location
		delete(p.numToLocation, val)
		p.nums[location] = p.nums[len(p.nums)-1]
		p.nums = p.nums[:len(p.nums)-1]
		return true
	}
}

func (p *RandomizedSet) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(p.nums))
	return p.nums[i]
}
