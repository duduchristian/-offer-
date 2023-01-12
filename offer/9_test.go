package offer

import (
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	nums := []int{10, 100, 100, 10, 5, 2, 6, 100}
	k := 100
	if NumSubarrayProductLessThanK(nums, k) != 9 {
		t.Fatal("Not pass")
	}
}
