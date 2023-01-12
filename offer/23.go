package offer

import (
	"algorithm/utils"
)

//面试题23：两个链表的第1个重合节点
//题目：输入两个单向链表，请问如何找出它们的第1个重合节点。例如，图4.5中的两个链表的第1个重合节点的值是4。

func GetIntersectionNode(headA, headB ListNode) ListNode {
	count1, count2 := countList(headA), countList(headB)
	delta := utils.Abs(count1 - count2)
	var longer, shorter ListNode
	if count1 > count2 {
		longer = headA
		shorter = headB
	} else {
		longer = headB
		shorter = headA
	}
	node1 := longer
	for i := 0; i < delta; i++ {
		node1 = node1.next
	}
	node2 := shorter
	for node1 != node2 {
		node1 = node1.next
		node2 = node2.next
	}
	return node1
}

func countList(head ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.next
	}
	return count
}
