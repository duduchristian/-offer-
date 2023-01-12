package offer

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	root := ConstructCBT([]int{5, 2, 4, 1, 6, 3, 7})
	if PathSum(root, 8) != 2 {
		t.Fatal("Error!")
	}

	if PathSum(root, 11) != 1 {
		t.Fatal("Error!!")
	}

	root = ConstructCBT([]int{5, 2, 4, 1, 6, 3, 7, 0, 0, 0, 0})
	if PathSum(root, 8) != 6 {
		t.Fatal("Error!!!")
	}
}
