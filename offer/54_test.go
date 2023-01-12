package offer

import "testing"

func TestConvertBST(t *testing.T) {
	root := ConstructCBT([]int{4, 2, 6, 1, 3, 5, 7})
	root2 := ConstructCBT([]int{22, 27, 13, 28, 25, 18, 7})
	root = ConvertBST(root)
	if !IsEqual(root, root2) {
		t.Fatal("error!!!")
	}
}
