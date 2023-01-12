package offer

//面试题29：排序的循环链表

func Insert(head Node, insertVal int) Node {
	node := NewNode(insertVal)
	if head == nil {
		head = node
		head.next = head
	} else if head.next == head {
		head.next = node
		node.next = head
	} else {
		insertCore(head, node)
	}
	return head
}

func insertCore(head, node Node) {
	cur := head
	next := head.next
	biggest := head
	for !(cur.val <= node.val && next.val >= node.val) && next != head {
		cur = next
		next = next.next
		if cur.val >= biggest.val {
			biggest = cur
		}
	}

	if cur.val <= node.val && next.val >= node.val {
		cur.next = node
		node.next = next
	} else {
		node.next = biggest.next
		biggest.next = node
	}
}
