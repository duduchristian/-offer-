package offer

import (
	"fmt"
	"strings"
)

type ListNode *struct {
	val  int
	next ListNode
}

func NewListNode(value int) ListNode {
	return &struct {
		val  int
		next ListNode
	}{value, nil}
}

func GenList(input []int) ListNode {
	dummy := NewListNode(0)
	cur := dummy
	for _, num := range input {
		cur.next = NewListNode(num)
		cur = cur.next
	}
	return dummy.next
}

func ToString(head ListNode) string {
	builder := strings.Builder{}
	for head != nil {
		builder.WriteString(fmt.Sprintf("%d|", head.val))
		head = head.next
	}
	return builder.String()
}
