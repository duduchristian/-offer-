package offer

import (
	"fmt"
	"testing"
)

func TestBSTIterator(t *testing.T) {
	root := ConstructCBT([]int{4, 2, 6, 1, 3, 5, 7})
	it := NewBSTIterator(root)
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
