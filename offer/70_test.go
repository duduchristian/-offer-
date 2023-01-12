package offer

import "testing"

func TestSingleNonDuplicate(t *testing.T) {
	if SingleNonDuplicate([]int{3}) != 3 {
		t.Fatal("Error!")
	}
	if SingleNonDuplicate([]int{1, 1, 2, 2, 3, 4, 4, 5, 5}) != 3 {
		t.Fatal("Error!!")
	}
	if SingleNonDuplicate([]int{1, 2, 2, 3, 3, 4, 4, 5, 5}) != 1 {
		t.Fatal("Error!!!")
	}
	if SingleNonDuplicate([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6}) != 6 {
		t.Fatal("Error!!!!")
	}
}
