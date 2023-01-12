package offer

import "testing"

func TestSumNumbers(t *testing.T) {
	root := ConstructCBT([]int{3, 9, 0, 5, 1, 2})
	if SumNumbers(root) != 1088 {
		t.Fatal("Error!")
	}
}
