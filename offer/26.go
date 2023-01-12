package offer

//面试题26：重排链表

func ReorderList(head ListNode) {
	dummy := NewListNode(0)
	dummy.next = head
	fast, slow := dummy, dummy
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next
		if fast.next != nil {
			fast = fast.next
		}
	}
	temp := slow.next
	slow.next = nil
	link(head, RevertList(temp), dummy)
}

func link(node1, node2, head ListNode) {
	prev := head
	for node1 != nil && node2 != nil {
		temp := node1.next
		prev.next = node1
		node1.next = node2
		prev = node2
		node1 = temp
		node2 = node2.next
	}

	if node1 != nil {
		prev.next = node1
	}
}
