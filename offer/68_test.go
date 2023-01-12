package offer

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	if SearchInsert([]int{1, 3, 6, 8}, 3) != 1 {
		t.Fatal("Error!")
	}
	if SearchInsert([]int{1, 3, 6, 8}, 5) != 2 {
		t.Fatal("Error!!")
	}
	if SearchInsert([]int{1, 3, 3, 3, 6, 8}, 3) != 1 {
		t.Fatal("Error!!!")
	}
	if SearchInsert([]int{1, 3, 3, 3, 6, 8}, 10) != 6 {
		t.Fatal("Error!!!!")
	}
}
