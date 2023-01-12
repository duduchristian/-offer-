package offer

//面试题21：删除倒数第k个节点
//题目：如果给定一个链表，请问如何删除链表中的倒数第k个节点？假设链表中节点的总数为n，那么1≤k≤n。要求只能遍历链表一次

func RemoveNthFromEnd(head ListNode, n int) ListNode {
	dummy := NewListNode(0)
	dummy.next = head
	left, right := dummy, head
	for i := 0; i < n; i++ {
		right = right.next
	}
	for right != nil {
		right = right.next
		left = left.next
	}
	left.next = left.next.next
	return dummy.next
}
