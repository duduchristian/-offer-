package offer

//面试题27：回文链表

func ListIsPalindrome(head ListNode) bool {
	if head == nil || head.next == nil {
		return true
	}

	slow, fast := head, head.next

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	secondHalf := slow.next
	slow.next = nil
	return equals(head, RevertList(secondHalf))
}

func equals(head1 ListNode, head2 ListNode) bool {
	for head1 != nil && head2 != nil {
		if head1.val != head2.val {
			return false
		}
		head1 = head1.next
		head2 = head2.next
	}
	return true
}
