package offer

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	list := GenList([]int{1, 2, 3, 4, 5})
	RemoveNthFromEnd(list, 2)
	if ToString(list) != "1|2|3|5|" {
		t.Fatal("error!")
	}
	RemoveNthFromEnd(list, 3)
	if ToString(list) != "1|3|5|" {
		t.Fatal("error!")
	}
}
