package offer

import (
	"algorithm/utils"
	"math"
)

//面试题66：单词之和

type MapSumNode *struct {
	value    int
	children map[rune]MapSumNode
}

func NewMapSumNode(value int) MapSumNode {
	return &struct {
		value    int
		children map[rune]MapSumNode
	}{value: value, children: map[rune]MapSumNode{}}
}

type MapSum struct {
	root MapSumNode
}

func NewMapSum() *MapSum {
	return &MapSum{
		root: NewMapSumNode(0),
	}
}

func (m *MapSum) Insert(word string, val int) {
	node := m.root
	for _, r := range word {
		if node.children[r] == nil {
			node.children[r] = NewMapSumNode(0)
		}
		node = node.children[r]
		node.value += val
	}
}

func (m *MapSum) Sum(prefix string) int {
	if len(prefix) == 0 {
		return 0
	}
	node := m.root
	ret := math.MaxInt
	for _, r := range prefix {
		if node.children[r] == nil {
			return 0
		}
		node = node.children[r]
		ret = utils.Min(ret, node.value)
	}
	return ret
}
