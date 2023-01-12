package offer

import (
	"algorithm/utils"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	res := TopKFrequent([]int{1, 2, 2, 1, 3, 1}, 2)
	if !utils.IntSlicesAreEqual(res, []int{2, 1}) {
		t.Fatal("Error!")
	}

	res = TopKFrequent([]int{1, 2, 2, 1, 3, 1, 3, 3, 3}, 2)
	if !utils.IntSlicesAreEqual(res, []int{1, 3}) {
		t.Fatal("Error!!")
	}

}
