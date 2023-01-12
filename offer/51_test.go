package offer

import (
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	root := ConstructCBT([]int{-9, 4, 20, 0, 0, 15, 7, 0, 0, 0, 0, -3})
	PruneTree(root)
	if MaxPathSum(root) != 42 {
		t.Fatal("Error!")
	}
	root = ConstructCBT([]int{-9, 4, 20, 0, 0, 15, 7, 0, 0, 0, 0, -3, 1})
	PruneTree(root)
	if MaxPathSum(root) != 43 {
		t.Fatal("Error!!")
	}
}
