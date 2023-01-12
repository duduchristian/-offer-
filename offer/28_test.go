package offer

import (
	"algorithm/utils"
	"testing"
)

func genNodeList(input []int) Node {
	if len(input) == 0 {
		return nil
	}
	head := NewNode(input[0])
	cur := head
	for i := 1; i < len(input); i++ {
		node := NewNode(input[i])
		cur.next = node
		cur = node
	}
	return head
}

func toIntSlice(head Node) []int {
	var ret []int
	for head != nil {
		ret = append(ret, head.val)
		head = head.next
	}
	return ret
}

func TestFlatten(t *testing.T) {
	head := genNodeList([]int{1, 2, 3, 7, 8, 9})
	head.next.next.child = genNodeList([]int{4, 6})
	head.next.next.child.child = genNodeList([]int{5})
	Flatten(head)
	if !utils.IntSlicesAreEqual([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, toIntSlice(head)) {
		t.Fatal("case 1 fail!!!")
	}
}
