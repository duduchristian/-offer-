package offer

//面试题22：链表中环的入口节点
//题目：如果一个链表中包含环，那么应该如何找出环的入口节点？
//从链表的头节点开始顺着next指针方向进入环的第1个节点为环的入口节点。例如，在如图4.3所示的链表中，环的入口节点是节点3。

func DetectCycle(head ListNode) ListNode {
	inLoop := getNodeInLoop(head)
	if inLoop == nil {
		return nil
	}
	node := head
	for node != inLoop {
		node = node.next
		inLoop = inLoop.next
	}
	return node
}

func DetectCycle1(head ListNode) ListNode {
	inLoop := getNodeInLoop(head)
	if inLoop == nil {
		return nil
	}
	loopCount := 1
	for n := inLoop; n.next != inLoop; n = n.next {
		loopCount++
	}
	slow, fast := head, head
	for i := 0; i < loopCount; i++ {
		fast = fast.next
	}
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}
	return slow
}

func getNodeInLoop(head ListNode) ListNode {
	if head == nil || head.next == nil {
		return nil
	}

	slow := head.next
	fast := slow.next
	for slow != nil && fast != nil {
		if slow == fast {
			return slow
		}
		slow = slow.next
		fast = fast.next
		if fast != nil {
			fast = fast.next
		}
	}
	return nil
}
